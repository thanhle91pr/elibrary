package repository

import (
	"context"
	"elibrary/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const SongsCollection = "songs"

type SongsRepository interface {
	CreateSongs(ctx context.Context, songs model.Songs) (string, error)
}

type songsRepository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewSongsRepository(db *mongo.Database) SongsRepository {
	return &songsRepository{
		col: db.Collection(SongsCollection),
	}
}

func (s *songsRepository) CreateSongs(ctx context.Context, songs model.Songs) (string, error) {
	songs.ID = primitive.NewObjectID().Hex()
	doc := bson.M{
		"_id":         songs.ID,
		"id":          songs.ID,
		"name":        songs.Name,
		"description": songs.Description,
	}
	_, err := s.col.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return songs.ID, err
}
