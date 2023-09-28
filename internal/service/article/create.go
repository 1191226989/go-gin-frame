package article

import "go-gin-frame/internal/model/article"

type CreateArticleData struct {
	Title   string
	Content string
}

// Create
func (s *service) Create(articleData *CreateArticleData) (uint, error) {
	m, err := article.NewModel()
	if err != nil {
		return 0, err
	}

	id, err := m.Create(&article.Article{
		Title:   articleData.Title,
		Content: articleData.Content,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}
