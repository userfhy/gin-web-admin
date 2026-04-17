package dictTypeService

import (
	model "gin-web-admin/app/models"
	"gin-web-admin/utils/gredis"
	"strings"
)

const (
	dictCachePrefix = "sys:dict:data:"
	dictCacheAllKey = "sys:dict:data:all"
)

type CachedDictItem struct {
	Label     string `json:"label"`
	Value     string `json:"value"`
	Status    int    `json:"status"`
	Sort      int    `json:"sort"`
	CSSClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault int    `json:"isDefault"`
	Remark    string `json:"remark"`
}

type CachedDictBucket struct {
	Name string           `json:"name"`
	Type string           `json:"type"`
	List []CachedDictItem `json:"list"`
}

func normalizeCacheDictType(dictType string) string {
	return strings.TrimSpace(strings.ToLower(dictType))
}

func dictCacheKey(dictType string) string {
	return dictCachePrefix + normalizeCacheDictType(dictType)
}

func loadDictCache(dictType string) ([]CachedDictItem, bool, error) {
	var cached []CachedDictItem
	found, err := gredis.GetJSON(dictCacheKey(dictType), &cached)
	return cached, found, err
}

func LoadCachedDictData(dictType string) ([]CachedDictItem, bool, error) {
	return loadDictCache(dictType)
}

func refreshDictCache(dictTypes ...string) error {
	if gredis.RedisConn == nil {
		return nil
	}

	status := 1
	typeList, err := model.GetAllDictTypes(&status)
	if err != nil {
		return err
	}
	dataList, err := model.GetAllDictData(&status)
	if err != nil {
		return err
	}

	typeNameMap := make(map[string]string, len(typeList))
	cacheMap := make(map[string][]CachedDictItem, len(typeList))
	for _, item := range typeList {
		typeNameMap[item.Type] = item.Name
		cacheMap[item.Type] = []CachedDictItem{}
	}

	for _, item := range dataList {
		dictType := normalizeCacheDictType(item.DictType)
		if _, ok := typeNameMap[dictType]; !ok {
			continue
		}
		cacheMap[dictType] = append(cacheMap[dictType], CachedDictItem{
			Label:     item.Label,
			Value:     item.Value,
			Status:    item.Status,
			Sort:      item.Sort,
			CSSClass:  item.CSSClass,
			ListClass: item.ListClass,
			IsDefault: item.IsDefault,
			Remark:    item.Remark,
		})
	}

	targets := make(map[string]struct{})
	if len(dictTypes) == 0 {
		for dictType := range typeNameMap {
			targets[dictType] = struct{}{}
		}
	} else {
		for _, dictType := range dictTypes {
			normalized := normalizeCacheDictType(dictType)
			if normalized == "" {
				continue
			}
			targets[normalized] = struct{}{}
		}
	}

	allBuckets := make([]CachedDictBucket, 0, len(targets))
	for dictType := range targets {
		items := cacheMap[dictType]
		if typeNameMap[dictType] == "" && items == nil {
			if _, err := gredis.Delete(dictCacheKey(dictType)); err != nil {
				return err
			}
			continue
		}

		bucket := CachedDictBucket{
			Name: typeNameMap[dictType],
			Type: dictType,
			List: items,
		}
		if err := gredis.SetJSON(dictCacheKey(dictType), bucket.List, 0); err != nil {
			return err
		}
		allBuckets = append(allBuckets, bucket)
	}

	if len(dictTypes) == 0 {
		if err := gredis.SetJSON(dictCacheAllKey, allBuckets, 0); err != nil {
			return err
		}
	}

	return nil
}
