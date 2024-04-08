package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/michelemendel/genvaeg/auth"
	repo "github.com/michelemendel/genvaeg/repository"
	"github.com/michelemendel/genvaeg/urlshortener"
)

type HandlerContext struct {
	Ctx          context.Context
	Session      *auth.Session
	Repo         *repo.Repo
	UrlShortener *urlshortener.URLShortener
}

func NewHandlerContext(echo *echo.Echo, session *auth.Session, repo *repo.Repo, us *urlshortener.URLShortener) *HandlerContext {
	return &HandlerContext{
		Ctx:          context.Background(),
		Session:      session,
		Repo:         repo,
		UrlShortener: us,
	}
}
