package order

// 创建
func (m *model) Create(order *Order) (uint, error) {
	tx := m.db.Create(&order)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return order.ID, nil
}
