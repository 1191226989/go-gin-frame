package order

// 详情
func (m *model) Detail(id int) (*Order, error) {
	order := Order{}
	tx := m.db.First(&order, id)
	if tx.Error != nil {
		return &order, tx.Error
	}
	return &order, nil
}
