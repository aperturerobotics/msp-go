package packet

import (
	"bytes"
	"strings"
)

// MspBoxNames contains checkbox names.
// Direction: FC->
// MessageID: 116
type MspBoxNames struct {
	BoxNames []string
}

// New builds a new instance of the packet.
func (p *MspBoxNames) New() Packet {
	return &MspBoxNames{}
}

// GetID returns this packet ID.
func (p *MspBoxNames) GetID() uint8 {
	return 116
}

// Marshal marshals the packet to a byte array.
func (p *MspBoxNames) Marshal() ([]byte, error) {
	return writeLSBBuf([]byte(strings.Join(p.BoxNames, ";")))
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspBoxNames) Unmarshal(data []byte) error {
	str := make([]byte, len(data))
	if err := readLSB(bytes.NewReader(data), &str); err != nil {
		return err
	}

	p.BoxNames = strings.Split(string(str), ";")
	return nil
}

func init() {
	registerPacketType(&MspBoxNames{})
}
