package article

// 文章详情
func (m *model) Detail(id uint) (*Article, error) {
	article := Article{}
	tx := m.db.First(&article, id)
	if tx.Error != nil {
		return &article, tx.Error
	}
	return &article, nil
}
