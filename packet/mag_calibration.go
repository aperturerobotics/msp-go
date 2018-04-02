package packet

import (
	"errors"
)

// MspMagCalibration commands the MSP calibrate the magnetometer.
// Direction: ->FC
// MessageID: 206
type MspMagCalibration struct{}

// New builds a new instance of the packet.
func (p *MspMagCalibration) New() Packet {
	return &MspMagCalibration{}
}

// GetID returns this packet ID.
func (p *MspMagCalibration) GetID() uint8 {
	return 206
}

// Marshal marshals the packet to a byte array.
func (p *MspMagCalibration) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspMagCalibration) Unmarshal(data []byte) error {
	if len(data) != 0 {
		return errors.New("unexpected body for MSP_ACC_CALIBRATION")
	}

	return nil
}

func init() {
	registerPacketType(&MspMagCalibration{})
}
