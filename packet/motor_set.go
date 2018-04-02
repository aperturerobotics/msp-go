package packet

import (
	"bytes"
	"errors"
)

// MspSetMotor sets individual motor value (with DYNBALANCE config).
// Direction: ->FC
// MessageID: 214
type MspSetMotor struct {
	// Motors contains the motor values between 1000 and 2000.
	Motors []uint16
}

// New builds a new instance of the packet.
func (p *MspSetMotor) New() Packet {
	return &MspSetMotor{}
}

// GetID returns this packet ID.
func (p *MspSetMotor) GetID() uint8 {
	return 214
}

// Marshal marshals the packet to a byte array.
func (p *MspSetMotor) Marshal() ([]byte, error) {
	if len(p.Motors) != MotorCount {
		if len(p.Motors) == 0 {
			p.Motors = make([]uint16, MotorCount)
		} else {
			return nil, errors.New("incorrect motor count")
		}
	}

	datas := make([]interface{}, MotorCount)
	for i := range p.Motors {
		datas[i] = &p.Motors[i]
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetMotor) Unmarshal(data []byte) error {
	motors := make([]uint16, MotorCount)
	datas := make([]interface{}, MotorCount)
	for i := range motors {
		datas[i] = &motors[i]
	}

	p.Motors = motors
	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspSetMotor{})
}
