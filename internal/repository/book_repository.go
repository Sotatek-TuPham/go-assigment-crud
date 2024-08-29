package repository

import (
	"github.com/hauchu1196/ecombase/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func (r *BookRepository) GetDB() *gorm.DB {
	return r.db
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *BookRepository) Update(book *models.Book) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(book).Error
}

func (r *BookRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&models.Book{}, id).Error
	})
}
