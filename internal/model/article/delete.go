package article

// 文章删除
func (m *model) Delete(id uint) error {
	article := Article{}
	tx := m.db.Delete(&article, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
