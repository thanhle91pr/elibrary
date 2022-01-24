package delivery

import (
	"context"
	"elibrary/pkg/model"
	"elibrary/pkg/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

type handle struct {
	usBooks  usecase.BooksUseCase
	usSongs  usecase.SongsUseCase
	ucLabels usecase.LabelsUseCase
}

func HttpHandel(e *echo.Echo, usBooks usecase.BooksUseCase, usSongs usecase.SongsUseCase, ucLabels usecase.LabelsUseCase) {
	handle := &handle{
		usBooks:  usBooks,
		usSongs:  usSongs,
		ucLabels: ucLabels,
	}
	e.GET("/labels", handle.GetListLabels)
	e.POST("/books/create", handle.CreateBooks)
	e.POST("/songs/create", handle.CreateSongs)
	e.GET("/labels/generate/:total", handle.Generate)
}
func (h handle) Generate(c echo.Context) error {
	total := c.Param("total")
	intTotal, err := strconv.ParseInt(total, 10, 32)
	if err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	err = h.ucLabels.GenericsLabels(context.Background(), int(intTotal))
	return model.ResponseSuccess(c, nil)
}

func (h *handle) CreateSongs(c echo.Context) error {
	var err error
	songs := model.Songs{}
	if err = c.Bind(&songs); err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	res, err := h.usSongs.CreateSongs(context.Background(), songs)
	if err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	return model.ResponseSuccess(c, res)
}

// GetListLabel will get list label
func (h *handle) GetListLabels(c echo.Context) error {
	var req model.GetListLabelsReq
	err := c.Bind(&req)
	if err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	res, err := h.ucLabels.GetListLabels(context.Background(), req.Page, req.Size)
	if err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	return model.ResponseSuccess(c, res)
}
func (h *handle) CreateBooks(c echo.Context) error {
	var err error
	book := model.Books{}
	if err = c.Bind(&book); err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	res, err := h.usBooks.CreateBook(context.Background(), book)
	if err != nil {
		log.Fatal(err)
		return model.ResponseWithError(c, err)
	}
	return model.ResponseSuccess(c, res)
}
