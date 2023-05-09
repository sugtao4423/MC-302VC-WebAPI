package enlrequest

import (
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/codelist"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/data"
)

func Make(serviceID byte, properties []data.Property) *data.Data {
	ps := len(properties)
	p := make([]data.Property, ps)
	for i := 0; i < ps; i++ {
		p[i] = data.Property{
			ID:          properties[i].ID,
			DataCounter: byte(len(properties[i].Value)),
			Value:       properties[i].Value,
		}
	}

	return &data.Data{
		Header: &data.Header{
			Protocol: 0x10,
			Format:   0x81,
		},
		TransactionID: []byte{0x00, 0x00},
		Source: &data.SEOJ{
			ClassGroupCode: codelist.Group_ProfileObject,
			ClassCode:      codelist.Class_NodeProfile,
			InstanceCode:   0x01,
		},
		Destination: &data.DEOJ{
			ClassGroupCode: codelist.Group_HousingFacilitiesRelatedDevice,
			ClassCode:      codelist.Class_HotWaterGenerator,
			InstanceCode:   0x01,
		},
		ServiceID:    serviceID,
		PropertySize: byte(ps),
		Properties:   p,
	}
}
