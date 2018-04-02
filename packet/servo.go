package packet

import (
	"bytes"
)

// MspServo contains servo data.
// Direction: FC->
// MessageID: 103
type MspServo struct {
	// Servos contains the servo values between 1000 and 2000.
	Servos []uint16
}

// New builds a new instance of the packet.
func (p *MspServo) New() Packet {
	return &MspServo{}
}

// GetID returns this packet ID.
func (p *MspServo) GetID() uint8 {
	return 103
}

// Marshal marshals the packet to a byte array.
func (p *MspServo) Marshal() ([]byte, error) {
	servos := make([]uint16, ServoCount)
	datas := make([]interface{}, ServoCount)
	for i := range servos {
		datas[i] = &servos[i]
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspServo) Unmarshal(data []byte) error {
	servos := make([]uint16, ServoCount)
	datas := make([]interface{}, ServoCount)
	for i := range servos {
		datas[i] = &servos[i]
	}

	p.Servos = servos
	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspServo{})
}
