package packet

import (
	"errors"
)

// MspAccCalibration commands the MSP calibrate the accel.
// Direction: ->FC
// MessageID: 205
type MspAccCalibration struct{}

// New builds a new instance of the packet.
func (p *MspAccCalibration) New() Packet {
	return &MspAccCalibration{}
}

// GetID returns this packet ID.
func (p *MspAccCalibration) GetID() uint8 {
	return 205
}

// Marshal marshals the packet to a byte array.
func (p *MspAccCalibration) Marshal() ([]byte, error) {
	return nil, nil
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspAccCalibration) Unmarshal(data []byte) error {
	if len(data) != 0 {
		return errors.New("unexpected body for MSP_ACC_CALIBRATION")
	}

	return nil
}

func init() {
	registerPacketType(&MspAccCalibration{})
}
