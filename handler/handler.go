package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	repo "github.com/michelemendel/genvaeg/repository"
)

type HandlerContext struct {
	Ctx  context.Context
	Repo *repo.Repo
}

func NewHandlerContext(echo *echo.Echo, repo *repo.Repo) *HandlerContext {
	return &HandlerContext{
		Ctx:  context.Background(),
		Repo: repo,
	}
}
