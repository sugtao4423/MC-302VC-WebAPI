package webapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *WebAPI) GetStatus(c echo.Context) error {
	status, err := a.mc302vc.GetStatus()
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, status)
	return nil
}
