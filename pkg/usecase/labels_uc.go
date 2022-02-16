package usecase

import (
	"context"
	"elibrary/pkg/model"
	"elibrary/pkg/repository"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type LabelsUseCase interface {
	GenericsLabels(ctx context.Context, totalLabels int) error
	GetListLabels(ctx context.Context, page int64, size int64) (lables []model.Label, err error)
	SetLabels(ctx context.Context, request model.SetLabelsRequest) error
	FindLabels(ctx context.Context, request model.FindLabelsRequest) []string
}

type labelsUseCase struct {
	labelsRepo repository.LabelsRepository
	combosRepo repository.CombosRepository
	songsRepo  repository.SongsRepository
	booksRepo  repository.BooksRepository
}

func NewLabelsUseCase(
	labelsRepo repository.LabelsRepository,
	combosRepo repository.CombosRepository,
	songsRepo repository.SongsRepository,
	booksRepo repository.BooksRepository) LabelsUseCase {
	return &labelsUseCase{
		labelsRepo: labelsRepo,
		combosRepo: combosRepo,
		songsRepo:  songsRepo,
		booksRepo:  booksRepo,
	}
}
func (l labelsUseCase) FindLabels(ctx context.Context, request model.FindLabelsRequest) []string {
	var labels []string

	if request.Song != "" {
		filter := bson.M{"name": request.Song}
		song := l.songsRepo.FindOne(ctx, filter, &options.FindOneOptions{})
		if song.Labels != "" {
			labels = append(labels, song.Labels)
		}
	}

	if request.Books != "" {
		filter := bson.M{"name": request.Books}
		books := l.booksRepo.FindOne(ctx, filter, &options.FindOneOptions{})
		if books.Labels != "" {
			labels = append(labels, books.Labels)
		}
	}

	if request.Song != "" && request.Books != "" {
		filter := bson.M{"books": request.Books, "songs": request.Song}
		combos := l.combosRepo.FindOne(ctx, filter, &options.FindOneOptions{})
		if combos.Labels != "" {
			labels = append(labels, combos.Labels)
		}
	}
	//remove duplicate
	if len(labels) > 1 {
		var result []string
		check := make(map[string]bool)
		for _, label := range labels {
			if _, ok := check[label]; ok {
				continue
			}
			result = append(result, label)
			check[label] = true
		}
		return result
	}
	return labels
}

func (l labelsUseCase) SetLabels(ctx context.Context, request model.SetLabelsRequest) error {
	if request.Label == "" {
		return errors.New("labels is not nil")
	}

	//set label to songs
	if request.Songs != "" {
		err := l.songsRepo.SetLabelToSongs(ctx, request.Label, request.Songs)
		if err != nil {
			log.Fatal(fmt.Sprintf("l.songsRepo.SetLabelToSongs error:%s", err))
		}
	}

	//set label to book
	if request.Book != "" {
		err := l.booksRepo.SetLabelToBooks(ctx, request.Label, request.Book)
		if err != nil {
			log.Fatal(fmt.Sprintf("l.songsRepo.SetLabelToBook error:%s", err))
		}
	}

	//add new combos
	if request.Songs != "" && request.Book != "" && request.Label != "" {
		err := l.combosRepo.UpsertCombos(ctx, model.Combos{Songs: request.Songs, Books: request.Book}, request.Label)
		if err != nil {
			log.Fatal(err)
		}

	}
	return nil
}

func (l labelsUseCase) GenericsLabels(ctx context.Context, totalLabels int) error {
	if totalLabels <= 0 {
		return errors.New("totalLabels not nil")
	}
	var docs []interface{}
	for i := 1; i <= totalLabels; i++ {
		//id := primitive.NewObjectID().Hex()
		docs = append(docs, bson.M{
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
