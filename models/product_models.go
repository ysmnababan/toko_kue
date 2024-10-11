package models

// Product represents the structure for your table
type Product struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID uint    `gorm:"not null" json:"category_id"` // Foreign key reference to a category table
	Name       string  `gorm:"size:255;not null" json:"name"`
	Code       string  `gorm:"size:50;not null;unique" json:"code"`
	Price      float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock      int     `gorm:"not null" json:"stock"`
}
