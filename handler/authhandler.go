package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michelemendel/genvaeg/entity"
	"github.com/michelemendel/genvaeg/util"
)

// curl -X POST -L -v 'http://localhost:8080/signup' -d 'name=abe&pw=abe&repeatpw=abe' -H 'Content-Type: application/x-www-form-urlencoded'
func (h *HandlerContext) SignupHandler(c echo.Context) error {
	username := c.FormValue("name")
	password := c.FormValue("pw")
	repeatPassword := c.FormValue("repeatpw")

	fmt.Printf("SignupHandler: name:%s, pw:%s, repeatpw:%s", username, password, repeatPassword)

	if password != repeatPassword {
		return c.String(http.StatusBadRequest, "Passwords do not match")
	}

	hashedPassword, _ := util.HashPassword(password)
	user := entity.NewUser(username, hashedPassword)
	err := h.Repo.CreateUser(user)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Failed to create user: %s\n", err))
	}

	return c.String(http.StatusOK, "User created\n")
}

// curl -X POST -L -v -c cookies.txt 'http://localhost:8080/login' -d 'name=abe&pw=abe' -H 'Content-Type: application/x-www-form-urlencoded'
func (h *HandlerContext) LoginHandler(c echo.Context) error {
	username := c.FormValue("name")
	password := c.FormValue("pw")

	user, err := h.Repo.GetUserByName(username)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Failed to get user: %s\n", err))
	}

	if !util.ValidatePassword(password, user.HashedPassword) {
		return c.String(http.StatusBadRequest, "Invalid password")
	}

	h.Session.Login(c, username)
	return c.String(http.StatusOK, "You are logged.\n")
}

// curl -L -v 'http://localhost:8080/logout'
func (h *HandlerContext) LogoutHandler(c echo.Context) error {
	h.Session.Logout(c)
	return c.String(http.StatusOK, "User logged out\n")
}
