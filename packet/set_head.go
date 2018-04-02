package packet

import (
	"bytes"
)

// MspSetHead sets a new heading lock.
// Direction: ->FC
// MessageID: 211
type MspSetHead struct {
	// Heading is the heading to go to, -180 to 180.
	Heading int16
}

// New builds a new instance of the packet.
func (p *MspSetHead) New() Packet {
	return &MspSetHead{}
}

// GetID returns this packet ID.
func (p *MspSetHead) GetID() uint8 {
	return 211
}

// Marshal marshals the packet to a byte array.
func (p *MspSetHead) Marshal() ([]byte, error) {
	return writeLSBBuf(&p.Heading)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetHead) Unmarshal(data []byte) error {
	return readLSB(bytes.NewReader(data), &p.Heading)
}

func init() {
	registerPacketType(&MspSetHead{})
}
