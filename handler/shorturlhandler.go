package handler

import (

	// "fmt"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michelemendel/genvaeg/entity"
	"github.com/michelemendel/genvaeg/util"
)

// curl -L -v http://localhost:8080
func (h *HandlerContext) NotYetImplementedHandler(c echo.Context) error {
	fmt.Println("NotYetImplementedHandler")
	return c.String(http.StatusOK, "This route is not yet implemented\n")
}

// curl -L -v http://localhost:8080/short
func (h *HandlerContext) RedirectHandler(c echo.Context) error {
	fmt.Println("RedirectHandler")
	// return c.Redirect(http.StatusMovedPermanently, "https://www.vg.no")
	return c.String(http.StatusOK, fmt.Sprintf("Redirecting to %s\n", "https://www.vg.no"))
}

// curl -L -v http://localhost:8080/create/http://vg.no
func (h *HandlerContext) CreateShortURLHandler(c echo.Context) error {
	fullURL := c.Param("full")

	if !h.UrlShortener.IsURLValid(fullURL) {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Invalid URL %s\n", fullURL))
	}

	baseURL := c.Scheme() + "://" + c.Request().Host
	hPw, _ := util.HashPassword("dummy")
	dummyUser := entity.NewUser("dummy", hPw)
	shortURL := h.UrlShortener.MakeShortURL(baseURL, dummyUser)

	fmt.Println("FullURL:", fullURL)
	fmt.Println("BaseURL:", baseURL)
	fmt.Println("ShortURL:", shortURL)

	return c.String(http.StatusOK, fmt.Sprintf("Short URL %s\n", shortURL))
}

// curl -L -v http://localhost:8080/list
func (h *HandlerContext) ListURLsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ListURLsHandler is not yet implemented\n")
}
