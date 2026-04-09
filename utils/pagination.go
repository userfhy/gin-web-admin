package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	List        any   `json:"list"`
	Total       int64 `json:"total"`
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`
}

// Pagination 表示分页请求参数
type Pagination struct {
	Page     int
	PageSize int
}

// Offset 计算当前页对应的偏移量
func (p Pagination) Offset() int {
	if p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

// Limit 返回分页条数
func (p Pagination) Limit() int {
	if p.PageSize <= 0 {
		return 0
	}
	return p.PageSize
}

// Result 根据分页信息快速生成 PageResult
func (p Pagination) Result(list any, total int64) PageResult {
	return PageResult{
		List:        list,
		Total:       total,
		PageSize:    p.PageSize,
		CurrentPage: p.Page,
	}
}

type paginationConfig struct {
	pageKeys    []string
	sizeKeys    []string
	defaultPage int
	defaultSize int
	maxSize     int
}

// PaginationOption 自定义分页解析行为
type PaginationOption func(*paginationConfig)

// WithDefaultPage 指定默认页码
func WithDefaultPage(page int) PaginationOption {
	return func(cfg *paginationConfig) {
		if page > 0 {
			cfg.defaultPage = page
		}
	}
}

// WithDefaultPageSize 指定默认每页数量
func WithDefaultPageSize(size int) PaginationOption {
	return func(cfg *paginationConfig) {
		if size > 0 {
			cfg.defaultSize = size
		}
	}
}

// WithMaxPageSize 设置最大页大小，传入 0 表示不限制
func WithMaxPageSize(size int) PaginationOption {
	return func(cfg *paginationConfig) {
		cfg.maxSize = size
	}
}

// WithParamAliases 为分页参数增加别名（优先级最高）
func WithParamAliases(pageKey, sizeKey string) PaginationOption {
	return func(cfg *paginationConfig) {
		if pageKey != "" {
			cfg.pageKeys = prependUnique(cfg.pageKeys, pageKey)
		}
		if sizeKey != "" {
			cfg.sizeKeys = prependUnique(cfg.sizeKeys, sizeKey)
		}
	}
}

func prependUnique(src []string, key string) []string {
	for _, k := range src {
		if k == key {
			return src
		}
	}
	return append([]string{key}, src...)
}

func defaultPaginationConfig() paginationConfig {
	return paginationConfig{
		pageKeys:    []string{"pageNum", "page", "p"},
		sizeKeys:    []string{"pageSize", "size", "n"},
		defaultPage: 1,
		defaultSize: 10,
		maxSize:     100,
	}
}

// GetPagination 统一解析分页参数（支持 pageNum/page/p 与 pageSize/size/n）
func GetPagination(c *gin.Context, opts ...PaginationOption) (Pagination, error) {
	cfg := defaultPaginationConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	page, err := pickIntParam(c, cfg.pageKeys)
	if err != nil {
		return Pagination{}, err
	}
	if page <= 0 {
		page = cfg.defaultPage
	}

	size, err := pickIntParam(c, cfg.sizeKeys)
	if err != nil {
		return Pagination{}, err
	}
	if size <= 0 {
		size = cfg.defaultSize
	}
	if cfg.maxSize > 0 && size > cfg.maxSize {
		size = cfg.maxSize
	}

	return Pagination{Page: page, PageSize: size}, nil
}

func pickIntParam(c *gin.Context, keys []string) (int, error) {
	for _, key := range keys {
		if value, ok := c.GetQuery(key); ok {
			v, err := strconv.Atoi(value)
			if err != nil {
				return 0, fmt.Errorf("%s 参数必须为数字", key)
			}
			return v, nil
		}
	}
	return 0, nil
}
