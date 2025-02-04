package models

type OrderItem struct {
	ID       int     `gorm:"primaryKey;autoIncrement"`
	OrderID  int     `gorm:"not null"`
	Order    Order   `gorm:"foreignKey:OrderID"`
	FishID   int     `gorm:"not null"`
	Fish     Fish    `gorm:"foreignKey:FishID"`
	Quantity int     `gorm:"not null"`
	Price    float64
}
