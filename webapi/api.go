package webapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc"
)

type WebAPI struct {
	mc302vc *mc302vc.MC302VC
}

type StatusPostBody struct {
	Status bool `json:"status"`
}

type TimerTimePostBody struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func New(mc302vc *mc302vc.MC302VC) *WebAPI {
	return &WebAPI{mc302vc}
}

func (a *WebAPI) ErrorHandler(err error, c echo.Context) {
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error": err.Error(),
	})
}

func (a *WebAPI) jsonSuccess(c echo.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": true,
	})
}
