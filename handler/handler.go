package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	repo "github.com/michelemendel/genvaeg/repository"
	"github.com/michelemendel/genvaeg/urlshortener"
)

type HandlerContext struct {
	Ctx          context.Context
	Repo         *repo.Repo
	UrlShortener *urlshortener.URLShortener
}

func NewHandlerContext(echo *echo.Echo, repo *repo.Repo, us *urlshortener.URLShortener) *HandlerContext {
	return &HandlerContext{
		Ctx:          context.Background(),
		Repo:         repo,
		UrlShortener: us,
	}
}
