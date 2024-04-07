package handler

import (

	// "fmt"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HandlerContext) NotYetImplementedHandler(c echo.Context) error {
	fmt.Println("NotYetImplementedHandler")
	return c.String(http.StatusOK, "This route is not yet implemented\n")
}

func (h *HandlerContext) RedirectHandler(c echo.Context) error {
	fmt.Println("RedirectHandler")

	// return c.Redirect(http.StatusMovedPermanently, "https://www.vg.no")
	return c.String(http.StatusOK, fmt.Sprintf("Redirecting to %s\n", "https://www.vg.no"))
}
