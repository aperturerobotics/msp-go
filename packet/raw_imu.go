package packet

import (
	"bytes"
)

// MspRawIMU contains raw IMU data.
// unit: it depends on ACC sensor and is based on ACC_1G definition
// Direction: FC->
// MessageID: 102
type MspRawIMU struct {
	// AccX is the acceleration in the X direction.
	AccX int16
	// AccY is the acceleration in the Y direction.
	AccY int16
	// AccZ is the acceleration in the Z direction.
	AccZ int16
	// GyrX is the gyro reading in the X direction.
	GyrX int16
	// GyrY is the gyro reading in the Y direction.
	GyrY int16
	// GyrZ is the gyro reading in the Z direction.
	GyrZ int16
	// MagX is the magnetometer reading in the X direction.
	MagX int16
	// MagY is the magnetometer reading in the Y direction.
	MagY int16
	// MagZ is the magnetometer reading in the Z direction.
	MagZ int16
}

// New builds a new instance of the packet.
func (p *MspRawIMU) New() Packet {
	return &MspRawIMU{}
}

// GetID returns this packet ID.
func (p *MspRawIMU) GetID() uint8 {
	return 102
}

// Marshal marshals the packet to a byte array.
func (p *MspRawIMU) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.AccX,
		&p.AccY,
		&p.AccZ,
		&p.GyrX,
		&p.GyrY,
		&p.GyrZ,
		&p.MagX,
		&p.MagY,
		&p.MagZ,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspRawIMU) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.AccX,
		&p.AccY,
		&p.AccZ,
		&p.GyrX,
		&p.GyrY,
		&p.GyrZ,
		&p.MagX,
		&p.MagY,
		&p.MagZ,
	)
}

func init() {
	registerPacketType(&MspRawIMU{})
}
