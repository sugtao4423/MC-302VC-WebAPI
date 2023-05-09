package mc302vc

import (
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/codelist"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/data"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/hwg"
)

type Status struct {
	OperationStatus             bool
	WaterTemperature            int
	BathTemperature             int
	BathAutoTimerStatus         bool
	BathAutoTimerTime           [2]int
	BathOperationStatus         bool
	BathAutoModeStatus          bool
	BathAdditionalHeatingStatus bool
}

func (m *MC302VC) GetStatus() (*Status, error) {
	s := byte(codelist.ESV_Get)
	p := []data.Property{
		{ID: hwg.EPC_OPERATION_STATUS},
		{ID: hwg.EPC_WATER_TEMPERATURE},
		{ID: hwg.EPC_BATH_TEMPERATURE_SETTING},
		{ID: hwg.EPC_BATH_AUTO_ON_TIMER_RESERVE_SETTING},
		{ID: hwg.EPC_BATH_AUTO_ON_TIMER_TIME_SETTING},
		{ID: hwg.EPC_BATH_OPERATION_STATUS_MONITORING},
		{ID: hwg.EPC_BATH_AUTO_MODE_SETTING},
		{ID: hwg.EPC_BATH_ADDITIONAL_HEATING_OPERATION_SETTING},
	}
	data, err := m.request(s, p)
	if err != nil {
		return nil, err
	}

	timerTime := [2]int{
		int(data.Properties[4].Value[0]),
		int(data.Properties[4].Value[1]),
	}

	return &Status{
		OperationStatus:             data.Properties[0].Value[0] == hwg.EDT_ON,
		WaterTemperature:            int(data.Properties[1].Value[0]),
		BathTemperature:             int(data.Properties[2].Value[0]),
		BathAutoTimerStatus:         data.Properties[3].Value[0] == hwg.EDT_YES,
		BathAutoTimerTime:           timerTime,
		BathOperationStatus:         data.Properties[5].Value[0] == hwg.EDT_YES,
		BathAutoModeStatus:          data.Properties[6].Value[0] == hwg.EDT_YES,
		BathAdditionalHeatingStatus: data.Properties[7].Value[0] == hwg.EDT_YES,
	}, nil
}
