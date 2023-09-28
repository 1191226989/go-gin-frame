package order

// 删除
func (m *model) Delete(id int) error {
	order := Order{}
	tx := m.db.Delete(&order, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
