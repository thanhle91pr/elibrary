package usecase

import (
	"context"
	"elibrary/pkg/model"
	"elibrary/pkg/repository"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type LabelsUseCase interface {
	GenericsLabels(ctx context.Context, totalLabels int) error
	GetListLabels(ctx context.Context, page int64, size int64) (lables []model.Label, err error)
}

type labelsUseCase struct {
	labelsRepo repository.LabelsRepository
}

func NewLabelsUseCase(labelsRepo repository.LabelsRepository) LabelsUseCase {
	return &labelsUseCase{labelsRepo: labelsRepo}
}

func (l labelsUseCase) GenericsLabels(ctx context.Context, totalLabels int) error {
	if totalLabels <= 0 {
		return errors.New("totalLabels not nil")
	}
	var docs []interface{}
	for i := 1; i <= totalLabels; i++ {
		id := primitive.NewObjectID().Hex()
		docs = append(docs, bson.M{
			"_id":         id,
			"id":          id,
			"name":        fmt.Sprintf("labels%v", i),
			"description": fmt.Sprintf("description labels%v", i),
		})
	}
	_, err := l.labelsRepo.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func (l labelsUseCase) GetListLabels(ctx context.Context, page int64, size int64) (lables []model.Label, err error) {
	if page <= 1 {
		page = 1
	}
	if size > 20 {
		size = 20
	}
	var skip, limit int64

	skip = (page - 1) * size
	limit = skip + size
	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.M{"_id": 1})
	lables, err = l.labelsRepo.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return lables, err
}
