package mc302vc

import (
	"net"
	"time"

	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/codelist"
	enldata "github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/data"
	"github.com/sugtao4423/MC-302VC-WebAPI/echonetlite/hwg"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc/enlrequest"
	"github.com/sugtao4423/MC-302VC-WebAPI/mc302vc/enlresponse"
)

type Client struct {
	conn *net.UDPConn
}

type MC302VC struct {
	addr   string
	client *Client
}

func New(mc302vcAddr string) (*MC302VC, error) {
	clientUdpAddr := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: 3610,
	}
	client, err := net.ListenUDP("udp", clientUdpAddr)

	if err != nil {
		return nil, err
	}

	return &MC302VC{
		addr:   mc302vcAddr,
		client: &Client{conn: client},
	}, nil
}

func (m *MC302VC) Close() error {
	return m.client.conn.Close()
}

func (c *Client) read() (*enldata.Data, error) {
	c.conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	buf := make([]byte, 1024)
	_, _, err := c.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}

	data := echonetlite.Parse(buf)
	return data, nil
}

func (m *MC302VC) request(serviceID byte, properties []enldata.Property) (*enldata.Data, error) {
	conn, err := net.Dial("udp", m.addr+":3610")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	d := enlrequest.Make(serviceID, properties)
	_, err = conn.Write(echonetlite.ToBytes(d))
	if err != nil {
		return nil, err
	}

	data, err := m.client.read()
	if err != nil {
		return nil, err
	}

	err = enlresponse.Validate(d, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *MC302VC) requestSetBool(epc byte, on bool) error {
	p := []enldata.Property{
		{
			ID:    epc,
			Value: []byte{hwg.EDT_YES},
		},
	}
	if !on {
		p[0].Value[0] = hwg.EDT_NO
	}

	_, err := m.request(codelist.ESV_SetC, p)
	if err != nil {
		return err
	}

	return nil
}
