package listing

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(listing *Listing) error
	Update(listing *Listing) error
	Delete(id uint) error
	FindByID(id uint) (*Listing, error)
	FindAll() ([]Listing, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(listing *Listing) error {
	return r.db.Create(listing).Error
}

func (r *repository) Update(listing *Listing) error {
	return r.db.Save(listing).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Listing{}, id).Error
}

func (r *repository) FindByID(id uint) (*Listing, error) {
	var listing Listing
	err := r.db.First(&listing, id).Error
	return &listing, err
}

func (r *repository) FindAll() ([]Listing, error) {
	var listings []Listing
	err := r.db.Find(&listings).Error
	return listings, err
}
