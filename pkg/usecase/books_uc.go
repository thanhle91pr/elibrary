package usecase

import (
	"context"
	"elibrary/pkg/model"
	"elibrary/pkg/repository"
)

type BooksUseCase interface {
	CreateBook(ctx context.Context, books model.Books) (string, error)
}

type bookUseCase struct {
	repoBooks repository.BooksRepository
}

func (b *bookUseCase) CreateBook(ctx context.Context, books model.Books) (string, error) {
	return b.repoBooks.CreateBooks(ctx, books)
}

func NewBookUseCase(repo repository.BooksRepository) BooksUseCase {
	return &bookUseCase{repoBooks: repo}
}
