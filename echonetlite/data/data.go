package data

type Header struct {
	Protocol byte
	Format   byte
}

type SEOJ struct {
	ClassGroupCode byte
	ClassCode      byte
	InstanceCode   byte
}

type DEOJ struct {
	ClassGroupCode byte
	ClassCode      byte
	InstanceCode   byte
}

type Property struct {
	ID          byte   // EPC
	DataCounter byte   // PDC
	Value       []byte // EDT
}

type Data struct {
	Header        *Header
	TransactionID []byte // TID
	Source        *SEOJ
	Destination   *DEOJ
	ServiceID     byte // ESV
	PropertySize  byte // OPC
	Properties    []Property
}
