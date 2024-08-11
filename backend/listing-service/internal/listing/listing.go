package listing

import "gorm.io/gorm"

type Listing struct {
	ID          uint     `gorm:"primaryKey"`
	Title       string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	Address     string   `gorm:"not null"`
	Price       float64  `gorm:"not null"`
	Images      []string `gorm:"type:text[]"`
}

func (l *Listing) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
