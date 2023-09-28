package article

import "go-gin-frame/internal/model/article"

type Service interface {
	Create(articleData *CreateArticleData) (uint, error)
	Delete(id uint) error
	Detail(id uint) (*article.Article, error)
	PageList(searchData *ListSearchData) ([]*article.Article, error)
	PageCount(searchData *ListSearchData) (int64, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
