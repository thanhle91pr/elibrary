package main

import (
	"context"
	"elibrary/config"
	"elibrary/pkg/delivery"
	"elibrary/pkg/repository"
	"elibrary/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//init cf env
	cf := config.InitConfig()
	//init db
	mgo := repository.NewMongoClient(cf.MongoUri)
	defer func() {
		_ = mgo.Disconnect(context.Background())
	}()
	db := mgo.Database("library")
	booksRepo := repository.NewBookRepository(db)
	songsRepo := repository.NewSongsRepository(db)
	labelsRepo := repository.NewLabelsRepository(db)
	combosRepo := repository.NewCombosRepository(db)
	booksUc := usecase.NewBookUseCase(booksRepo)
	songsUc := usecase.NewSongsUseCase(songsRepo)
	labelsUc := usecase.NewLabelsUseCase(labelsRepo, combosRepo, songsRepo, booksRepo)
	delivery.HttpHandel(e, booksUc, songsUc, labelsUc)
	e.Logger.Fatal(e.Start(":8088"))
}
