package packet

import (
	"bytes"
)

// MspMotor contains motor data.
// Direction: FC->
// MessageID: 104
type MspMotor struct {
	// Motors contains the motor values between 1000 and 2000.
	Motors [MotorCount]uint16
}

// New builds a new instance of the packet.
func (p *MspMotor) New() Packet {
	return &MspMotor{}
}

// GetID returns this packet ID.
func (p *MspMotor) GetID() uint8 {
	return 103
}

// Marshal marshals the packet to a byte array.
func (p *MspMotor) Marshal() ([]byte, error) {
	motors := make([]uint16, MotorCount)
	datas := make([]interface{}, MotorCount)
	for i := range motors {
		datas[i] = &motors[i]
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspMotor) Unmarshal(data []byte) error {
	motors := make([]uint16, MotorCount)
	datas := make([]interface{}, MotorCount)
	for i := range motors {
		datas[i] = &motors[i]
	}

	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspMotor{})
}
