package usecase

import (
	"context"
	"elibrary/pkg/model"
	"elibrary/pkg/repository"
	"log"
)

type SongsUseCase interface {
	CreateSongs(ctx context.Context, songs model.Songs) (string, error)
}
type songsUseCase struct {
	songsRepo repository.SongsRepository
}

func (s *songsUseCase) CreateSongs(ctx context.Context, songs model.Songs) (string, error) {
	res, err := s.songsRepo.CreateSongs(ctx, songs)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return res, err
}

func NewSongsUseCase(songsRepository repository.SongsRepository) SongsUseCase {
	return &songsUseCase{
		songsRepo: songsRepository,
	}
}
