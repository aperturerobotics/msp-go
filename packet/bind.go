package packet

import (
	"errors"
)

// MspBind commands the FC bind spektrum sattelite.
// Direction: ->FC
// MessageID: 240
type MspBind struct{}

// New builds a new instance of the packet.
func (p *MspBind) New() Packet {
	return &MspBind{}
}

// GetID returns this packet ID.
func (p *MspBind) GetID() uint8 {
	return 240
}

// Marshal marshals the packet to a byte array.
func (p *MspBind) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspBind) Unmarshal(data []byte) error {
	if len(data) != 0 {
		return errors.New("unexpected body for MSP_ACC_CALIBRATION")
	}

	return nil
}

func init() {
	registerPacketType(&MspBind{})
}
