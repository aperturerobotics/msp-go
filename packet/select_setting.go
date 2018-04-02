package packet

import (
	"bytes"
)

// MspSelectSetting selects a setting configuration (there are multiple).
// Direction: ->FC
// MessageID: 210
type MspSelectSetting struct {
	// Setting is the setting ID to use.
	Setting uint8
}

// New builds a new instance of the packet.
func (p *MspSelectSetting) New() Packet {
	return &MspSelectSetting{}
}

// GetID returns this packet ID.
func (p *MspSelectSetting) GetID() uint8 {
	return 210
}

// Marshal marshals the packet to a byte array.
func (p *MspSelectSetting) Marshal() ([]byte, error) {
	return writeLSBBuf(&p.Setting)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSelectSetting) Unmarshal(data []byte) error {
	return readLSB(bytes.NewReader(data), &p.Setting)
}

func init() {
	registerPacketType(&MspSelectSetting{})
}
