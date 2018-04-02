package packet

import (
	"bytes"
	"strings"
)

// MspPidNames contains pid names.
// Direction: FC->
// MessageID: 117
type MspPidNames struct {
	PidNames []string
}

// New builds a new instance of the packet.
func (p *MspPidNames) New() Packet {
	return &MspPidNames{}
}

// GetID returns this packet ID.
func (p *MspPidNames) GetID() uint8 {
	return 117
}

// Marshal marshals the packet to a byte array.
func (p *MspPidNames) Marshal() ([]byte, error) {
	return writeLSBBuf([]byte(strings.Join(p.PidNames, ";")))
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspPidNames) Unmarshal(data []byte) error {
	str := make([]byte, len(data))
	if err := readLSB(bytes.NewReader(data), &str); err != nil {
		return err
	}

	p.PidNames = strings.Split(string(str), ";")
	return nil
}

func init() {
	registerPacketType(&MspPidNames{})
}
