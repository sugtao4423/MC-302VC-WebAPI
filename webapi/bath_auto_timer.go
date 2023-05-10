package webapi

import (
	"github.com/labstack/echo/v4"
)

func (a *WebAPI) SetBathAutoTimer(c echo.Context) error {
	var body StatusPostBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	err := a.mc302vc.SetBathAutoTimer(body.Status)
	if err != nil {
		return err
	}

	a.jsonSuccess(c)
	return nil
}

func (a *WebAPI) SetBathAutoTimerTime(c echo.Context) error {
	var body TimerTimePostBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	err := a.mc302vc.SetBathAutoTimerTime(body.Hour, body.Minute)
	if err != nil {
		return err
	}

	a.jsonSuccess(c)
	return nil
}
