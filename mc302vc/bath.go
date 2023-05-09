package mc302vc

import (
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/hwg"
)

func (m *MC302VC) SetBathAuto(on bool) error {
	return m.requestSetBool(hwg.EPC_BATH_AUTO_MODE_SETTING, on)
}

func (m *MC302VC) SetBathAdditionalHeating(on bool) error {
	return m.requestSetBool(hwg.EPC_BATH_ADDITIONAL_HEATING_OPERATION_SETTING, on)
}
