package packet

import (
	"bytes"
)

// MspAttitude contains attitude data from the FC.
// Direction: FC->
// MessageID: 108
type MspAttitude struct {
	// AngX is the X angle, -1800 to 1800, unit 1/10 degree.
	AngX int16
	// AngY is the Y angle, -900 to 900, unit 1/10 degree.
	AngY int16
	// Heading is the heading direction, -180 to 180 compass degrees.
	Heading int16
}

// New builds a new instance of the packet.
func (p *MspAttitude) New() Packet {
	return &MspAttitude{}
}

// GetID returns this packet ID.
func (p *MspAttitude) GetID() uint8 {
	return 108
}

// Marshal marshals the packet to a byte array.
func (p *MspAttitude) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.AngX,
		&p.AngY,
		&p.Heading,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspAttitude) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.AngX,
		&p.AngY,
		&p.Heading,
	)
}

func init() {
	registerPacketType(&MspAttitude{})
}
