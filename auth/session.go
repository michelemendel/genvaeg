package auth

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	consts "github.com/michelemendel/genvaeg/constants"
	repo "github.com/michelemendel/genvaeg/repository"
	"github.com/michelemendel/genvaeg/util"
)

type TokenType string
type UserSession struct {
	Name  string
	Token TokenType
}
type Session struct {
	LoggedInUsers map[TokenType]UserSession
	Repo          *repo.Repo
}

func NewSession(repo *repo.Repo) *Session {
	return &Session{
		LoggedInUsers: make(map[TokenType]UserSession),
		Repo:          repo,
	}
}

func (s *Session) Login(c echo.Context, username string) {
	token := s.SetSession(c)
	s.LoggedInUsers[TokenType(token)] = UserSession{Name: username, Token: TokenType(token)}
}

func (s *Session) SetSession(c echo.Context) string {
	token := util.GenerateUUID()
	sess, _ := session.Get(consts.AUTH_SESSION_NAME, c)
	sess.Options = &sessions.Options{
		Path: "/",
		// MaxAge: 30, // 30 seconds
		MaxAge: 3600 * 5, // 5 minutes
		// MaxAge:   86400 * 1, // 1 day
		HttpOnly: true,
	}
	sess.Values[consts.AUTH_TOKEN_NAME] = token
	sess.Save(c.Request(), c.Response())
	return token
}

func (s *Session) Logout(c echo.Context) error {
	sess, _ := session.Get(consts.AUTH_SESSION_NAME, c)
	token := sess.Values[consts.AUTH_TOKEN_NAME]
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	if token != nil {
		delete(s.LoggedInUsers, TokenType(token.(string)))
	}
	return nil
}

func (s *Session) GetLoggedInUser(c echo.Context) (UserSession, error) {
	sess, _ := session.Get(consts.AUTH_SESSION_NAME, c)
	token := sess.Values[consts.AUTH_TOKEN_NAME]
	if token != nil {
		user := s.LoggedInUsers[TokenType(token.(string))]
		return user, nil
	}
	return UserSession{}, fmt.Errorf("no user found")
}

func (s *Session) GetLoggedInUsers() map[TokenType]UserSession {
	return s.LoggedInUsers
}
