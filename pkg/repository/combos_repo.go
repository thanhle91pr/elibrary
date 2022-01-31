package repository

import (
	"context"
	"elibrary/pkg/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CombosRepository interface {
	UpsertCombos(ctx context.Context, combos model.Combos, labelsName string) error
}
type combosRepository struct {
	db *mongo.Database
}

func (c combosRepository) UpsertCombos(ctx context.Context, combos model.Combos, labelsName string) error {
	col := c.db.Collection("combos")
	filter := bson.M{"books": combos.Books, "songs": combos.Songs}
	opts := options.Update().SetUpsert(true)
	update := bson.M{
		"$set": bson.M{
			"label": labelsName,
			"songs": combos.Songs,
			"books": combos.Books,
		},
	}
	_, err := col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Errorf(fmt.Sprintf("col.UpdateOne error %s", err))
	}
	return  err
}
func NewCombosRepository(db *mongo.Database) CombosRepository {
	return &combosRepository{db: db}
}
