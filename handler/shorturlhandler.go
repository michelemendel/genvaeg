package handler

import (

	// "fmt"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// curl -L -v 'http://localhost:8080/r/<short code>'
func (h *HandlerContext) RedirectHandler(c echo.Context) error {
	shortURLPath := c.Param("shorturlpath")
	fullURL, err := h.Repo.GetFullURLByShortURL(shortURLPath)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Short URL not found %s\n", shortURLPath))
	}

	return c.Redirect(http.StatusMovedPermanently, fullURL)
}

// curl -L -v -b cookies.txt 'http://localhost:8080/create?url=http://vg.no'
// Bad url
// curl -L -v -b cookies.txt 'http://localhost:8080/create?url=xxxxxx'
// Bad user
// curl -L -v -b cookies.txt 'http://localhost:8080/create?url=http://vg.no'
func (h *HandlerContext) CreateShortURLHandler(c echo.Context) error {
	fullURL := c.QueryParam("url")

	userSession, err := h.Session.GetLoggedInUser(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	userName := userSession.Name
	fmt.Println("[CreateShortURLHandler]: userName from session", userName)
	user, err := h.Repo.GetUserByName(userName)
	fmt.Println("[CreateShortURLHandler]: user from db", user)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("User not found %s\n", userName))
	}

	if !h.UrlShortener.IsURLValid(fullURL) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid URL %s\n", fullURL))
	}

	err = h.UrlShortener.MakeShortURL(fullURL, user)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Failed to make short URL: %s\n", err))
	}
	return c.String(http.StatusOK, "Short URL created\n")
}

// curl -L -v 'http://localhost:8080/urls'
func (h *HandlerContext) ListURLsHandler(c echo.Context) error {
	urls, err := h.Repo.GetAllURLPairs()
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Failed to get all URL pairs: %s\n", err))
	}

	urlsAsJson, err := json.Marshal(urls)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Failed to marshal URL pairs: %s\n", err))
	}

	return c.JSON(http.StatusOK, string(urlsAsJson))
}
