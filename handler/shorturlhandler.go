package handler

import (

	// "fmt"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: remove before take off
// curl -L -v 'http://localhost:8080'
func (h *HandlerContext) NotYetImplementedHandler(c echo.Context) error {
	fmt.Println("NotYetImplementedHandler")
	return c.String(http.StatusOK, "This route is not yet implemented\n")
}

// curl -L -v 'http://localhost:8080/<short code>'
func (h *HandlerContext) RedirectHandler(c echo.Context) error {
	shortURLPath := c.Param("shorturlpath")
	fullURL, err := h.Repo.GetFullURLByShortURL(shortURLPath)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Short URL not found %s\n", shortURLPath))
	}

	return c.Redirect(http.StatusMovedPermanently, fullURL)
}

// curl -L -v 'http://localhost:8080/create?url=http://vg.no&uuid=424c8ccf-648d-4826-8f39-687f4a16d8e9'
// Bad url
// curl -L -v 'http://localhost:8080/create?url=xxxxxx&uuid=424c8ccf-648d-4826-8f39-687f4a16d8e9'
// Bad user
// curl -L -v 'http://localhost:8080/create?url=http://vg.no&uuid=424c8ccf-648d-4826-8f39-687f4a16d8e9'
func (h *HandlerContext) CreateShortURLHandler(c echo.Context) error {
	fullURL := c.QueryParam("url")
	userUUID := c.QueryParam("uuid")

	user, err := h.Repo.GetUserByUUID(userUUID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("User not found %s\n", userUUID))
	}

	if !h.UrlShortener.IsURLValid(fullURL) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid URL %s\n", fullURL))
	}

	shortURL := h.UrlShortener.MakeShortURL(fullURL, user)
	return c.String(http.StatusOK, fmt.Sprintf("Short URL created, %s\n", shortURL))
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
