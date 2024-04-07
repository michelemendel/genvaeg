package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/michelemendel/genvaeg/constants"
	"github.com/michelemendel/genvaeg/handler"
)

func Routes(e *echo.Echo, hCtx *handler.HandlerContext) {
	// Auth
	e.POST(constants.ROUTE_SIGNUP, hCtx.NotYetImplementedHandler)
	e.POST(constants.ROUTE_LOGIN, hCtx.NotYetImplementedHandler)
	e.GET(constants.ROUTE_LOGOUT, hCtx.NotYetImplementedHandler)

	// Shorten
	e.GET(constants.ROUTE_CREATE_SHORT_URL, hCtx.CreateShortURLHandler)
	e.GET(constants.ROUTE_LIST_URLS, hCtx.ListURLsHandler)
	e.GET(constants.ROUTE_REDIRECT+":shorturlpath", hCtx.RedirectHandler)
}
