package webapi

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc"
)

type WebAPI struct {
	mc302vc  *mc302vc.MC302VC
	username string
	password string
}

type StatusPostBody struct {
	Status *bool `json:"status" validate:"required"`
}

type TimerTimePostBody struct {
	Hour   *int `json:"hour"   validate:"required"`
	Minute *int `json:"minute" validate:"required"`
}

func New(mc302vc *mc302vc.MC302VC, username, password string) *WebAPI {
	return &WebAPI{mc302vc, username, password}
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

func (a *WebAPI) BasicAuthMiddleware() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		return username == a.username && password == a.password, nil
	})
}

type JsonValidator struct {
	validator *validator.Validate
}

func (a *WebAPI) NewJsonValidator() *JsonValidator {
	return &JsonValidator{validator: validator.New()}
}

func (cv *JsonValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
