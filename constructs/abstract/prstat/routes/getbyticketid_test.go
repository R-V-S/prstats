package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	// Setup Echo
	e := echo.New()
	// Following Add can hopefully be removed at some point in the future, but is currently due to an open issue.
	// https://github.com/labstack/echo/issues/1492
	e.Add("GET", "/:arbitrary", func(echo.Context) error { return nil })
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("ticketid")
	c.SetParamValues("TB-1234")

	// Execute
	getByID := GetByTicketID()

	// Assertions
	if assert.NoError(t, getByID(c)) {
		fmt.Println(rec)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Regexp(t, `WIP!!!`, rec.Body.String()) // @TODO: define
	}
}
