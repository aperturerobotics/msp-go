package packet

import (
	"errors"
)

// MspEepRomWrite commands the FC write the settings to the EEPROM.
// Direction: ->FC
// MessageID: 250
type MspEepRomWrite struct{}

// New builds a new instance of the packet.
func (p *MspEepRomWrite) New() Packet {
	return &MspEepRomWrite{}
}

// GetID returns this packet ID.
func (p *MspEepRomWrite) GetID() uint8 {
	return 250
}

// Marshal marshals the packet to a byte array.
func (p *MspEepRomWrite) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspEepRomWrite) Unmarshal(data []byte) error {
	if len(data) != 0 {
		return errors.New("unexpected body for MSP_ACC_CALIBRATION")
	}

	return nil
}

func init() {
	registerPacketType(&MspEepRomWrite{})
}
