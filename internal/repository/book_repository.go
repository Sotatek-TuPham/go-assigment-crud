package repository

import (
	"github.com/hauchu1196/ecombase/internal/api/dto"
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

func (r *BookRepository) Update(id uint, payload dto.UpdateBookDTO) (*models.Book, error) {
	var book models.Book
	r.db.Where("id = ?", id).First(&book)
	if payload.BookName != nil {
		book.BookName = *payload.BookName
	}
	if payload.Author != nil {
		book.Author = *payload.Author
	}
	if payload.Category != nil {
		book.Category = *payload.Category
	}
	if payload.Shell != nil {
		book.Shell = *payload.Shell
	}
	err := r.db.Save(&book).Error
	return &book, err
}

func (r *BookRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&models.Book{}, id).Error
	})
}
