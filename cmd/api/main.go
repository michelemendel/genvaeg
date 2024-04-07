package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michelemendel/genvaeg/auth"
	consts "github.com/michelemendel/genvaeg/constants"
	"github.com/michelemendel/genvaeg/handler"
	repo "github.com/michelemendel/genvaeg/repository"
	"github.com/michelemendel/genvaeg/routes"
	"github.com/michelemendel/genvaeg/util"
)

func init() {
	util.InitEnv()
}

func main() {
	env := os.Getenv(consts.ENV_APP_ENV_KEY)
	webServerPort := os.Getenv(consts.ENV_WEB_SERVER_PORT_KEY)

	fmt.Printf("ENVIRONMENT:\nmode:%s\nwebServerPort:%s\n", env, webServerPort)

	r := repo.NewRepo()
	defer r.DB.Close()
	s := auth.NewSession(r)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv(consts.ENV_SESSION_KEY_KEY)))))
	e.Use(s.Authenticate)
	e.HTTPErrorHandler = customHTTPErrorHandler

	hCtx := handler.NewHandlerContext(e, r)
	routes.Routes(e, hCtx)
	slog.Debug("Starting server", "port", webServerPort)
	e.Logger.Fatal(e.Start(":" + webServerPort))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	httpErr, ok := err.(*echo.HTTPError)
	if ok {
		code = httpErr.Code
	}

	slog.Warn("httpError", "code", code)

	// Authentication (used forbidden here (403), since we use echo.ErrUnauthorized (401) for authorization)
	if code == echo.ErrForbidden.Code {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	// errorPage := fmt.Sprintf("./public/%d.html", code)
	// fileErr := c.File(errorPage)
	// if fileErr != nil {
	// 	c.Logger().Error(fileErr)
	// }

	c.String(code, fmt.Sprintf("Error %s\n", err.Error()))
}
