package services

import (
	"context"

	"github.com/ibanezv/fly-devs-practice_go/internal/models"
	"github.com/ibanezv/fly-devs-practice_go/internal/repository"
)

type BookService struct {
	repository repository.BooksRepository
}

func NewBookService(repository repository.BooksRepository) *BookService {
	return &BookService{repository: repository}
}

func (b BookService) GetById(ctx context.Context, id int64) (models.Book, error) {
	book, err := b.repository.Get(ctx, id)
	if err != nil {
		return models.Book{}, err
	}
	return repositoryToBook(ctx, *book), nil
}
