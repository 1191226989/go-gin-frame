package article

import "go-gin-frame/internal/model/article"

// Delete
func (s *service) Delete(id uint) error {
	m, err := article.NewModel()
	if err != nil {
		return err
	}

	err = m.Delete(id)

	return err
}
