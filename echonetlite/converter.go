package echonetlite

import (
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/data"
)

func Parse(buf []byte) *data.Data {
	header := &data.Header{
		Protocol: buf[0],
		Format:   buf[1],
	}

	transactionId := buf[2:4]

	dataSrc := &data.SEOJ{
		ClassGroupCode: buf[4],
		ClassCode:      buf[5],
		InstanceCode:   buf[6],
	}
	dataDst := &data.DEOJ{
		ClassGroupCode: buf[7],
		ClassCode:      buf[8],
		InstanceCode:   buf[9],
	}

	serviceId := buf[10]
	propertySize := buf[11]

	properties := make([]data.Property, propertySize)
	lastPos := 12
	for i := 0; i < int(propertySize); i++ {
		properties[i] = data.Property{
			ID:          buf[lastPos],
			DataCounter: buf[lastPos+1],
			Value:       buf[lastPos+2 : lastPos+2+int(buf[lastPos+1])],
		}
		lastPos += 2 + int(buf[lastPos+1])
	}

	return &data.Data{
		Header:        header,
		TransactionID: transactionId,
		Source:        dataSrc,
		Destination:   dataDst,
		ServiceID:     serviceId,
		PropertySize:  propertySize,
		Properties:    properties,
	}
}

func ToBytes(data *data.Data) []byte {
	buf := make([]byte, 0)
	buf = append(buf, data.Header.Protocol)
	buf = append(buf, data.Header.Format)
	buf = append(buf, data.TransactionID...)
	buf = append(buf, data.Source.ClassGroupCode)
	buf = append(buf, data.Source.ClassCode)
	buf = append(buf, data.Source.InstanceCode)
	buf = append(buf, data.Destination.ClassGroupCode)
	buf = append(buf, data.Destination.ClassCode)
	buf = append(buf, data.Destination.InstanceCode)
	buf = append(buf, data.ServiceID)
	buf = append(buf, data.PropertySize)
	for _, property := range data.Properties {
		buf = append(buf, property.ID)
		buf = append(buf, property.DataCounter)
		buf = append(buf, property.Value...)
	}
	return buf
}
