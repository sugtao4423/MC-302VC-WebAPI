package enlresponse

import (
	"fmt"

	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/codelist"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/data"
)

func Validate(request *data.Data, response *data.Data) error {
	if request.ServiceID == codelist.ESV_Get &&
		response.ServiceID != codelist.ESV_Get_Res {
		return fmt.Errorf("ServiceID is not ESV_Get_Res")
	}
	if (request.ServiceID == codelist.ESV_SetC || request.ServiceID == codelist.ESV_SetI) &&
		response.ServiceID != codelist.ESV_Set_Res {
		return fmt.Errorf("ServiceID is not ESV_Set_Res")
	}

	if len(request.Properties) != int(response.PropertySize) ||
		len(request.Properties) != len(response.Properties) {
		return fmt.Errorf("PropertySize is not match")
	}

	return nil
}
