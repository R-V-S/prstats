package routes

import (
	"github.com/labstack/echo/v4"
)

// Routes ...
func Routes(echo *echo.Echo) {
	prstats := echo.Group("/prstats")

	prstats.GET("/ticket-id/:ticketid", GetByTicketID())
}
