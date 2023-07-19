package services

import (
	"context"

	"github.com/ibanezv/fly-devs-practice_go/internal/models"
	"github.com/ibanezv/fly-devs-practice_go/internal/repository"
)

func repositoryToBook(ctx context.Context, book repository.Book) models.Book {
	return models.Book{Id: book.Id, Title: book.Title, Author: book.Author}
}
