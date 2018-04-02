package packet

import (
	"bytes"
)

// MspBoxIds contains checkbox ids.
// Direction: FC->
// MessageID: 119
type MspBoxIds struct {
	BoxIds []uint8
}

// New builds a new instance of the packet.
func (p *MspBoxIds) New() Packet {
	return &MspBoxIds{}
}

// GetID returns this packet ID.
func (p *MspBoxIds) GetID() uint8 {
	return 119
}

// Marshal marshals the packet to a byte array.
func (p *MspBoxIds) Marshal() ([]byte, error) {
	items := make([]interface{}, len(p.BoxIds))
	for i := range p.BoxIds {
		items[i] = &p.BoxIds[i]
	}

	return writeLSBBuf(items...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspBoxIds) Unmarshal(data []byte) error {
	p.BoxIds = make([]uint8, len(data))
	items := make([]interface{}, len(data))
	for i := range p.BoxIds {
		items[i] = &p.BoxIds[i]
	}

	return readLSB(bytes.NewReader(data), items...)
}

func init() {
	registerPacketType(&MspBoxIds{})
}
