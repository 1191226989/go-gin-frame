package article

import (
	"fmt"
)

type SearchData struct {
	Page     int // 第几页
	PageSize int // 每页数量
	Title    string
}

// 分页数据列表
func (m *model) PageList(searchData *SearchData) ([]*Article, error) {
	page := searchData.Page
	if page == 0 {
		page = 1
	}
	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	tx := m.db.Limit(pageSize).Offset(offset)
	if searchData.Title != "" {
		tx = tx.Where("title LIKE ?", fmt.Sprintf("%%%s%%", searchData.Title))
	}

	listData := make([]*Article, 0)
	result := tx.Find(&listData)

	return listData, result.Error
}

// 分页数据数量
func (m *model) PageCount(searchData *SearchData) (int64, error) {
	var count int64
	tx := m.db.Model(&Article{})
	if searchData.Title != "" {
		tx = tx.Where("title LIKE ?", fmt.Sprintf("%%%s%%", searchData.Title))
	}

	result := tx.Count(&count)
	return count, result.Error
}
