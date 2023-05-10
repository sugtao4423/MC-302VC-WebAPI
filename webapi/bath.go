package webapi

import (
	"github.com/labstack/echo/v4"
)

func (a *WebAPI) SetBathAuto(c echo.Context) error {
	var body StatusPostBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	err := a.mc302vc.SetBathAuto(body.Status)
	if err != nil {
		return err
	}

	a.jsonSuccess(c)
	return nil
}

func (a *WebAPI) SetBathAdditionalHeating(c echo.Context) error {
	var body StatusPostBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	err := a.mc302vc.SetBathAdditionalHeating(body.Status)
	if err != nil {
		return err
	}

	a.jsonSuccess(c)
	return nil
}
