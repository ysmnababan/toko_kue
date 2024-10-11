package models

type Category struct {
	CategoryID  uint   `gorm:"primaryKey;autoIncrement" json:"category_id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Code        string `gorm:"size:100;not null;unique" json:"code"`
	Description string `gorm:"type:text" json:"description"`
}
