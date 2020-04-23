package routes

import (
	"database/sql"
	"net/http"

	"github.com/R-V-S/prstats/constructs/abstract/prstat"
	"github.com/R-V-S/prstats/constructs/peripheral/errormessage"
	"github.com/R-V-S/prstats/constructs/symbiotes/github"
	"github.com/labstack/echo/v4"
)

// GetByTicketID ...
// @tags PRStat
// @Router /prstats/id/:id [get]
// @Summary Get an existing cognition by its ID
// @Produce  json
// @Param ID path string true "Ticket ID"
// @Success 200 {object} prstat.Aggregate
// @Failure 400 {object} errormessage.ErrorMessage
// @Failure 404 {object} errormessage.ErrorMessage
// @Failure 500 {object} errormessage.ErrorMessage
func GetByTicketID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		ticketid := c.Param("ticketid")

		prStatAggregate := prstat.Aggregate{}
		err = github.FetchPRStats(&prStatAggregate, ticketid)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, nil)
			}
			return c.JSON(http.StatusInternalServerError, errormessage.NewError("SYMBIOTE", err.Error()))
		}

		return c.JSON(http.StatusOK, prStatAggregate)
	}
}
