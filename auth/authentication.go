package auth

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/michelemendel/genvaeg/constants"
)

func (s *Session) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if os.Getenv(constants.ENV_BYPASS_LOGIN) == "true" {
			return next(c)
		}

		sess, _ := s.GetLoggedInUser(c)
		path := c.Path()

		if path == "/login" {
			return next(c)
		}

		if sess.Name == "" {
			return echo.ErrForbidden
		} else {
			return next(c)
		}
	}
}
