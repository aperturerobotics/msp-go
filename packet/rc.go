package packet

import (
	"bytes"
	"errors"
)

// MspRc contains RC data from the FC.
// ROLL/PITCH/YAW/THROTTLE/AUX1/AUX2/AUX3AUX4
// Direction: FC->
// MessageID: 105
type MspRc struct {
	// Channels contains the RC values between 1000 and 2000.
	Channels []uint16
}

// New builds a new instance of the packet.
func (p *MspRc) New() Packet {
	return &MspRc{}
}

// GetID returns this packet ID.
func (p *MspRc) GetID() uint8 {
	return 105
}

// Marshal marshals the packet to a byte array.
func (p *MspRc) Marshal() ([]byte, error) {
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
func (p *MspRc) Unmarshal(data []byte) error {
	channels := make([]uint16, ChannelCount)
	datas := make([]interface{}, ChannelCount)
	for i := range channels {
		datas[i] = &channels[i]
	}

	p.Channels = channels
	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspRc{})
}
