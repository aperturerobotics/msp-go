package packet

import (
	"errors"
)

// MspResetConf commands the FC reset all configuration to default.
// Direction: ->FC
// MessageID: 208
type MspResetConf struct{}

// New builds a new instance of the packet.
func (p *MspResetConf) New() Packet {
	return &MspResetConf{}
}

// GetID returns this packet ID.
func (p *MspResetConf) GetID() uint8 {
	return 208
}

// Marshal marshals the packet to a byte array.
func (p *MspResetConf) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspResetConf) Unmarshal(data []byte) error {
	if len(data) != 0 {
		return errors.New("unexpected body for MSP_ACC_CALIBRATION")
	}

	return nil
}

func init() {
	registerPacketType(&MspResetConf{})
}
