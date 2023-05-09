package mc302vc

import (
	"fmt"

	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/codelist"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/data"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/hwg"
)

func (m *MC302VC) SetBathAutoTimer(on bool) error {
	return m.requestSetBool(hwg.EPC_BATH_AUTO_ON_TIMER_RESERVE_SETTING, on)
}

func (m *MC302VC) SetBathAutoTimerTime(hour int, min int) error {
	if hour < 0 || hour > 23 {
		return fmt.Errorf("hour must be 0-23")
	}
	if min < 0 || min > 59 {
		return fmt.Errorf("min must be 0-59")
	}

	p := []data.Property{
		{
			ID:    hwg.EPC_BATH_AUTO_ON_TIMER_TIME_SETTING,
			Value: []byte{byte(hour), byte(min)},
		},
	}

	_, err := m.request(codelist.ESV_SetC, p)
	if err != nil {
		return err
	}

	return nil
}
