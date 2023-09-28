package article

import "go-gin-frame/internal/model/article"

type ListSearchData struct {
	Page     int // 第几页
	PageSize int // 每页数量
	Title    string
}

func (s *service) PageList(searchData *ListSearchData) ([]*article.Article, error) {
	m, err := article.NewModel()
	if err != nil {
		return nil, err
	}

	search := article.SearchData{
		Page:     searchData.Page,
		PageSize: searchData.PageSize,
		Title:    searchData.Title,
	}
	listData, err := m.PageList(&search)

	return listData, err
}

func (s *service) PageCount(searchData *ListSearchData) (int64, error) {
	m, err := article.NewModel()
	if err != nil {
		return 0, err
	}

	search := article.SearchData{
		Page:     searchData.Page,
		PageSize: searchData.PageSize,
		Title:    searchData.Title,
	}
	count, err := m.PageCount(&search)

	return count, err
}
