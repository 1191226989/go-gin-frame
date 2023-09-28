package article

// 文章创建
func (m *model) Create(article *Article) (uint, error) {
	tx := m.db.Create(&article)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return article.ID, nil
}
