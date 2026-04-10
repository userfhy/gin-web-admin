package sitePublicService

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	model "gin-web-admin/app/models"
	"gin-web-admin/internal/data"
	"gin-web-admin/utils"
	"gin-web-admin/utils/gredis"
	"gin-web-admin/utils/logging"
)

type PublicContentQuery struct {
	Pagination   utils.Pagination
	Keyword      string
	CategorySlug string
	TagSlug      string
}

type PublicCategoryVO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	ContentCount int64  `json:"contentCount"`
}

type PublicContentListVO struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	Slug          string             `json:"slug"`
	Summary       string             `json:"summary"`
	Cover         string             `json:"cover"`
	UpdatedAt     any                `json:"updatedAt"`
	PublishedAt   any                `json:"publishedAt"`
	ReadingMinute int                `json:"readingMinute"`
	Categories    []PublicCategoryVO `json:"categories"`
	Tags          []PublicTagVO      `json:"tags"`
}

type PublicContentDetailVO struct {
	ID             int                `json:"id"`
	Title          string             `json:"title"`
	Slug           string             `json:"slug"`
	Summary        string             `json:"summary"`
	Cover          string             `json:"cover"`
	Content        string             `json:"content"`
	SeoKeywords    string             `json:"seoKeywords"`
	SeoDescription string             `json:"seoDescription"`
	UpdatedAt      any                `json:"updatedAt"`
	PublishedAt    any                `json:"publishedAt"`
	ReadingMinute  int                `json:"readingMinute"`
	Categories     []PublicCategoryVO `json:"categories"`
	Tags           []PublicTagVO      `json:"tags"`
}

type PublicTagVO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Service struct {
	store *data.Store
}

var defaultService *Service

const (
	categoriesCacheKey = "site:categories:all"
	tagsCacheKey       = "site:tags:all"
	siteCacheTTL       = 5 * time.Minute
	detailCachePrefix  = "site:content:"
	listCachePrefix    = "site:content:list:"
)

func NewService(store *data.Store) *Service {
	return &Service{store: store}
}

func SetDefaultService(s *Service) {
	defaultService = s
}

func serviceInstance() *Service {
	if defaultService == nil {
		panic("site_public service not initialized")
	}
	return defaultService
}

func GetPublicTags() ([]PublicTagVO, error) {
	return serviceInstance().GetPublicTags()
}

func (s *Service) GetPublicTags() ([]PublicTagVO, error) {
	var cached []PublicTagVO
	if ok, err := gredis.GetJSON(tagsCacheKey, &cached); err == nil && ok {
		return cached, nil
	}

	status := 1
	tags, err := model.GetAllSiteTags(&status)
	if err != nil {
		return nil, err
	}
	result := make([]PublicTagVO, 0, len(tags))
	for _, item := range tags {
		result = append(result, PublicTagVO{
			ID:   item.ID,
			Name: item.Name,
			Slug: item.Slug,
		})
	}
	gredis.SetJSONAsync(tagsCacheKey, result, siteCacheTTL)
	return result, nil
}

func GetPublicCategories() ([]PublicCategoryVO, error) {
	return serviceInstance().GetPublicCategories()
}

func (s *Service) GetPublicCategories() ([]PublicCategoryVO, error) {
	var cached []PublicCategoryVO
	if ok, err := gredis.GetJSON(categoriesCacheKey, &cached); err == nil && ok {
		return cached, nil
	}

	status := 1
	categories, err := model.GetAllSiteCategories(&status)
	if err != nil {
		return nil, err
	}
	ids := make([]int, 0, len(categories))
	for _, item := range categories {
		ids = append(ids, item.ID)
	}
	countMap, err := model.CountSiteContentByCategoryIDs(ids)
	if err != nil {
		return nil, err
	}
	result := make([]PublicCategoryVO, 0, len(categories))
	for _, item := range categories {
		result = append(result, PublicCategoryVO{
			ID:           item.ID,
			Name:         item.Name,
			Slug:         item.Slug,
			Description:  item.Description,
			ContentCount: countMap[item.ID],
		})
	}
	gredis.SetJSONAsync(categoriesCacheKey, result, siteCacheTTL)
	return result, nil
}

func GetPublicContentList(query PublicContentQuery) (utils.PageResult, error) {
	return serviceInstance().GetPublicContentList(query)
}

func (s *Service) GetPublicContentList(query PublicContentQuery) (utils.PageResult, error) {
	cacheKey := listCacheKey(query)
	if pageResult, ok := s.getCachedContentList(cacheKey); ok {
		return pageResult, nil
	}

	pg := query.Pagination.Clone()
	var categoryID *int
	if slug := strings.TrimSpace(query.CategorySlug); slug != "" {
		category, err := model.GetSiteCategoryBySlug(slug)
		if err != nil {
			return utils.PageResult{}, err
		}
		if category == nil || category.Status != 1 {
			return pg.Result([]PublicContentListVO{}, 0), nil
		}
		categoryID = &category.ID
	}

	var tagID *int
	if slug := strings.TrimSpace(query.TagSlug); slug != "" {
		tag, err := model.GetSiteTagBySlug(slug)
		if err != nil {
			return utils.PageResult{}, err
		}
		if tag == nil || tag.Status != 1 {
			return pg.Result([]PublicContentListVO{}, 0), nil
		}
		tagID = &tag.ID
	}

	list, total, err := model.GetPublishedSiteContentList(pg, query.Keyword, categoryID, tagID)
	if err != nil {
		return utils.PageResult{}, err
	}

	categoryByContent, err := s.getPublicCategoriesByContentIDs(extractContentIDs(list))
	if err != nil {
		return utils.PageResult{}, err
	}
	tagByContent, err := s.getPublicTagsByContentIDs(extractContentIDs(list))
	if err != nil {
		return utils.PageResult{}, err
	}

	result := make([]PublicContentListVO, 0, len(list))
	for _, item := range list {
		result = append(result, PublicContentListVO{
			ID:            item.ID,
			Title:         item.Title,
			Slug:          item.Slug,
			Summary:       buildListSummary(item.Summary, item.Content),
			Cover:         item.Cover,
			UpdatedAt:     item.UpdatedAt,
			PublishedAt:   item.PublishedAt,
			ReadingMinute: readingMinute(item.Content),
			Categories:    categoryByContent[item.ID],
			Tags:          tagByContent[item.ID],
		})
	}

	pageResult := pg.Result(result, total)
	s.setCachedContentList(cacheKey, pageResult)
	return pageResult, nil
}

func GetPublicContentDetailBySlug(slug string) (*PublicContentDetailVO, error) {
	return serviceInstance().GetPublicContentDetailBySlug(slug)
}

func (s *Service) GetPublicContentDetailBySlug(slug string) (*PublicContentDetailVO, error) {
	if strings.TrimSpace(slug) == "" {
		return nil, fmt.Errorf("slug is required")
	}
	if detail, ok := s.getCachedContentDetail(slugCacheKey(slug)); ok {
		return detail, nil
	}

	item, err := model.GetSiteContentBySlug(slug)
	if err != nil {
		return nil, err
	}
	if item == nil || item.Status != 1 {
		return nil, nil
	}

	categoryByContent, err := s.getPublicCategoriesByContentIDs([]int{item.ID})
	if err != nil {
		return nil, err
	}
	tagByContent, err := s.getPublicTagsByContentIDs([]int{item.ID})
	if err != nil {
		return nil, err
	}

	detail := &PublicContentDetailVO{
		ID:             item.ID,
		Title:          item.Title,
		Slug:           item.Slug,
		Summary:        buildListSummary(item.Summary, item.Content),
		Cover:          item.Cover,
		Content:        item.Content,
		SeoKeywords:    item.SeoKeywords,
		SeoDescription: item.SeoDescription,
		UpdatedAt:      item.UpdatedAt,
		PublishedAt:    item.PublishedAt,
		ReadingMinute:  readingMinute(item.Content),
		Categories:     categoryByContent[item.ID],
		Tags:           tagByContent[item.ID],
	}
	s.setCachedContentDetail(slugCacheKey(slug), detail)
	return detail, nil
}

func GetPublicContentDetailByID(id int) (*PublicContentDetailVO, error) {
	return serviceInstance().GetPublicContentDetailByID(id)
}

func (s *Service) GetPublicContentDetailByID(id int) (*PublicContentDetailVO, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id is required")
	}
	cacheKey := idCacheKey(id)
	if detail, ok := s.getCachedContentDetail(cacheKey); ok {
		return detail, nil
	}

	item, err := model.GetSiteContentByID(id)
	if err != nil {
		return nil, err
	}
	if item == nil || item.Status != 1 {
		return nil, nil
	}

	categoryByContent, err := s.getPublicCategoriesByContentIDs([]int{item.ID})
	if err != nil {
		return nil, err
	}
	tagByContent, err := s.getPublicTagsByContentIDs([]int{item.ID})
	if err != nil {
		return nil, err
	}

	detail := &PublicContentDetailVO{
		ID:             item.ID,
		Title:          item.Title,
		Slug:           item.Slug,
		Summary:        buildListSummary(item.Summary, item.Content),
		Cover:          item.Cover,
		Content:        item.Content,
		SeoKeywords:    item.SeoKeywords,
		SeoDescription: item.SeoDescription,
		UpdatedAt:      item.UpdatedAt,
		PublishedAt:    item.PublishedAt,
		ReadingMinute:  readingMinute(item.Content),
		Categories:     categoryByContent[item.ID],
		Tags:           tagByContent[item.ID],
	}
	s.setCachedContentDetail(cacheKey, detail)
	return detail, nil
}

func (s *Service) getPublicTagsByContentIDs(contentIDs []int) (map[int][]PublicTagVO, error) {
	result := make(map[int][]PublicTagVO, len(contentIDs))
	if len(contentIDs) == 0 {
		return result, nil
	}
	tagIDsMap, err := model.GetTagIDsByContentIDs(contentIDs)
	if err != nil {
		return nil, err
	}
	allIDSet := map[int]struct{}{}
	for _, ids := range tagIDsMap {
		for _, id := range ids {
			allIDSet[id] = struct{}{}
		}
	}
	allIDs := make([]int, 0, len(allIDSet))
	for id := range allIDSet {
		allIDs = append(allIDs, id)
	}

	tagMap := map[int]PublicTagVO{}
	if len(allIDs) > 0 {
		tags, err := model.GetSiteTagsByIDs(allIDs)
		if err != nil {
			return nil, err
		}
		for _, t := range tags {
			if t.Status != 1 {
				continue
			}
			tagMap[t.ID] = PublicTagVO{
				ID:   t.ID,
				Name: t.Name,
				Slug: t.Slug,
			}
		}
	}

	for cid, ids := range tagIDsMap {
		items := make([]PublicTagVO, 0, len(ids))
		for _, id := range ids {
			if item, ok := tagMap[id]; ok {
				items = append(items, item)
			}
		}
		result[cid] = items
	}
	return result, nil
}

func (s *Service) getPublicCategoriesByContentIDs(contentIDs []int) (map[int][]PublicCategoryVO, error) {
	result := make(map[int][]PublicCategoryVO, len(contentIDs))
	if len(contentIDs) == 0 {
		return result, nil
	}
	categoryIDsMap, err := model.GetCategoryIDsByContentIDs(contentIDs)
	if err != nil {
		return nil, err
	}
	allIDSet := map[int]struct{}{}
	for _, ids := range categoryIDsMap {
		for _, id := range ids {
			allIDSet[id] = struct{}{}
		}
	}
	allIDs := make([]int, 0, len(allIDSet))
	for id := range allIDSet {
		allIDs = append(allIDs, id)
	}

	categoryMap := map[int]PublicCategoryVO{}
	if len(allIDs) > 0 {
		categories, err := model.GetSiteCategoriesByIDs(allIDs)
		if err != nil {
			return nil, err
		}
		for _, c := range categories {
			if c.Status != 1 {
				continue
			}
			categoryMap[c.ID] = PublicCategoryVO{
				ID:          c.ID,
				Name:        c.Name,
				Slug:        c.Slug,
				Description: c.Description,
			}
		}
	}

	for cid, ids := range categoryIDsMap {
		items := make([]PublicCategoryVO, 0, len(ids))
		for _, id := range ids {
			if item, ok := categoryMap[id]; ok {
				items = append(items, item)
			}
		}
		result[cid] = items
	}
	return result, nil
}

func extractContentIDs(list []*model.SiteContent) []int {
	ids := make([]int, 0, len(list))
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	return ids
}

func autoSummary(summary string, content string) string {
	if s := strings.TrimSpace(summary); s != "" {
		return s
	}
	plain := strings.TrimSpace(strings.ReplaceAll(content, "\n", " "))
	const maxLen = 140
	if utf8.RuneCountInString(plain) <= maxLen {
		return plain
	}
	runes := []rune(plain)
	return string(runes[:maxLen]) + "..."
}

func buildListSummary(summary string, content string) string {
	s := strings.TrimSpace(summary)
	if s == "" {
		s = strings.TrimSpace(strings.ReplaceAll(content, "\n", " "))
	}
	if s == "" {
		return ""
	}
	const maxLen = 140
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	runes := []rune(s)
	return string(runes[:maxLen]) + "..."
}

func readingMinute(content string) int {
	count := utf8.RuneCountInString(strings.TrimSpace(content))
	if count <= 0 {
		return 1
	}
	// 中文和英文混排按每分钟约400字估算
	minute := count / 400
	if count%400 != 0 {
		minute++
	}
	if minute <= 0 {
		return 1
	}
	return minute
}

func (s *Service) getCachedContentDetail(key string) (*PublicContentDetailVO, bool) {
	var detail PublicContentDetailVO
	ok, err := gredis.GetJSON(key, &detail)
	if err != nil {
		logging.Warnf("get content detail cache failed: %v", err)
		return nil, false
	}
	if !ok {
		return nil, false
	}
	return &detail, true
}

func (s *Service) setCachedContentDetail(key string, detail *PublicContentDetailVO) {
	gredis.SetJSONAsync(key, detail, siteCacheTTL)
}

func slugCacheKey(slug string) string {
	return detailCachePrefix + "slug:" + slug
}

func idCacheKey(id int) string {
	return detailCachePrefix + "id:" + fmt.Sprintf("%d", id)
}

func (s *Service) getCachedContentList(key string) (utils.PageResult, bool) {
	var result utils.PageResult
	ok, err := gredis.GetJSON(key, &result)
	if err != nil {
		logging.Warnf("get content list cache failed: %v", err)
		return utils.PageResult{}, false
	}
	return result, ok
}

func (s *Service) setCachedContentList(key string, result utils.PageResult) {
	gredis.SetJSONAsync(key, result, siteCacheTTL)
}

func listCacheKey(query PublicContentQuery) string {
	return fmt.Sprintf("%skeyword=%s:cat=%s:tag=%s:page=%d:size=%d",
		listCachePrefix,
		strings.TrimSpace(query.Keyword),
		strings.TrimSpace(query.CategorySlug),
		strings.TrimSpace(query.TagSlug),
		query.Pagination.Page,
		query.Pagination.PageSize,
	)
}

// InvalidatePublicCategories 清理分类相关缓存
func InvalidatePublicCategories() {
	gredis.DeleteKeysAsync(categoriesCacheKey)
	gredis.DeleteByPrefixAsync(listCachePrefix)
}

// InvalidatePublicTags 清理标签相关缓存
func InvalidatePublicTags() {
	gredis.DeleteKeysAsync(tagsCacheKey)
	gredis.DeleteByPrefixAsync(listCachePrefix)
}

// InvalidatePublicContent 清理内容列表与详情缓存
func InvalidatePublicContent(id int, slug string) {
	keys := make([]string, 0, 2)
	if id > 0 {
		keys = append(keys, idCacheKey(id))
	}
	if slug = strings.TrimSpace(slug); slug != "" {
		keys = append(keys, slugCacheKey(slug))
	}
	if len(keys) > 0 {
		gredis.DeleteKeysAsync(keys...)
	}
	gredis.DeleteByPrefixAsync(listCachePrefix)
}

// InvalidateAllPublicContent 完整清理所有内容缓存
func InvalidateAllPublicContent() {
	gredis.DeleteByPrefixAsync(detailCachePrefix)
	gredis.DeleteByPrefixAsync(listCachePrefix)
}
