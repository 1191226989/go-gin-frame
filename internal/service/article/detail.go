package article

import "go-gin-frame/internal/model/article"

// Detail
func (s *service) Detail(id uint) (*article.Article, error) {
	m, err := article.NewModel()
	if err != nil {
		return nil, err
	}

	article, err := m.Detail(id)

	return article, err
}
