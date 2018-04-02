package packet

import (
	"bytes"
)

// MspMotorPins contains motor pin data from the FC.
// Direction: FC->
// MessageID: 115
type MspMotorPins struct {
	Pins [MotorPinCount]uint8
}

// New builds a new instance of the packet.
func (p *MspMotorPins) New() Packet {
	return &MspMotorPins{}
}

// GetID returns this packet ID.
func (p *MspMotorPins) GetID() uint8 {
	return 115
}

// Marshal marshals the packet to a byte array.
func (p *MspMotorPins) Marshal() ([]byte, error) {
	items := make([]interface{}, MotorPinCount)
	for i := range items {
		items[i] = &p.Pins[i]
	}

	return writeLSBBuf(
		items...,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspMotorPins) Unmarshal(data []byte) error {
	items := make([]interface{}, MotorPinCount)
	for i := range items {
		items[i] = &p.Pins[i]
	}

	return readLSB(
		bytes.NewReader(data),

		items...,
	)
}

func init() {
	registerPacketType(&MspMotorPins{})
}
