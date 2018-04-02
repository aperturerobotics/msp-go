package packet

import (
	"bytes"
)

// MspAltitude contains altitude data from the FC.
// Direction: FC->
// MessageID: 109
type MspAltitude struct {
	// EstAlt is the estimated altitude in centimeters.
	EstAlt int32
	// Vario is the variance of the reading in cm/s.
	Vario int16
}

// New builds a new instance of the packet.
func (p *MspAltitude) New() Packet {
	return &MspAltitude{}
}

// GetID returns this packet ID.
func (p *MspAltitude) GetID() uint8 {
	return 109
}

// Marshal marshals the packet to a byte array.
func (p *MspAltitude) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.EstAlt,
		&p.Vario,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspAltitude) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.EstAlt,
		&p.Vario,
	)
}

func init() {
	registerPacketType(&MspAltitude{})
}
