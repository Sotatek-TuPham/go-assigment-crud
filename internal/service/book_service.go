package service

import (
	"github.com/hauchu1196/ecombase/internal/api/dto"
	"github.com/hauchu1196/ecombase/internal/models"
	"github.com/hauchu1196/ecombase/internal/repository"
)

type BookServiceInterface interface {
	CreateBook(book *models.Book) error
	GetBook(id uint) (*models.Book, error)
	UpdateBook(id uint, bookPayload dto.UpdateBookDTO) (*models.Book, error)
	DeleteBook(id uint) error
}

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(order *models.Book) error {
	return s.repo.Create(order)
}

func (s *BookService) GetBook(id uint) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) UpdateBook(id uint, bookPayload dto.UpdateBookDTO) (*models.Book, error) {
	return s.repo.Update(id, bookPayload)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}
