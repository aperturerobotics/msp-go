package packet

import (
	"bytes"
)

// MspStatus contains autopilot status data.
// Direction: FC->
// MessageID: 101
type MspStatus struct {
	// CycleTime is the cycle time in microseconds.
	CycleTime uint16
	// I2CErrorsCount contains the number of I2C errors observed.
	I2CErrorsCount uint16
	// Sensor indicates which sensors are active.
	// BARO<<1|MAG<<2|GPS<<3|SONAR<<4
	Sensor uint16
	// Flag contains a bit variable to indicate which BOX are active.
	// The bit position of each box depends on which are configured.
	Flag uint32
	// CurrentConfSet contains the current configuration setting.
	CurrentConfSet uint8
}

// New builds a new instance of the packet.
func (p *MspStatus) New() Packet {
	return &MspStatus{}
}

// GetID returns this packet ID.
func (p *MspStatus) GetID() uint8 {
	return 101
}

// Marshal marshals the packet to a byte array.
func (p *MspStatus) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.CycleTime,
		&p.I2CErrorsCount,
		&p.Sensor,
		&p.Flag,
		&p.CurrentConfSet,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspStatus) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.CycleTime,
		&p.I2CErrorsCount,
		&p.Sensor,
		&p.Flag,
		&p.CurrentConfSet,
	)
}

func init() {
	registerPacketType(&MspStatus{})
}
