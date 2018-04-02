package packet

import (
	"bytes"
	"errors"
)

// MspRcSet contains RC data to the FC.
// ROLL/PITCH/YAW/THROTTLE/AUX1/AUX2/AUX3AUX4
// This request is used to inject RC channel via MSP.
// Each chan overrides legacy RX as long as it is refreshed at least every second.
// Direction: ->FC
// MessageID: 200
type MspRcSet struct {
	// Channels contains the RC values between 1000 and 2000.
	Channels []uint16
}

// New builds a new instance of the packet.
func (p *MspRcSet) New() Packet {
	return &MspRcSet{}
}

// GetID returns this packet ID.
func (p *MspRcSet) GetID() uint8 {
	return 200
}

// Marshal marshals the packet to a byte array.
func (p *MspRcSet) Marshal() ([]byte, error) {
	if len(p.Channels) != ChannelCount {
		if len(p.Channels) == 0 {
			p.Channels = make([]uint16, ChannelCount)
		} else {
			return nil, errors.New("incorrect channel count")
		}
	}

	datas := make([]interface{}, ChannelCount)
	for i := range p.Channels {
		datas[i] = &p.Channels[i]
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspRcSet) Unmarshal(data []byte) error {
	channels := make([]uint16, ChannelCount)
	datas := make([]interface{}, ChannelCount)
	for i := range channels {
		datas[i] = &channels[i]
	}

	p.Channels = channels
	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspRcSet{})
}
