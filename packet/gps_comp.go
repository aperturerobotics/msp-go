package packet

import (
	"bytes"
)

// MspCompGps contains GPS distance to home data from the FC.
// Direction: FC->
// MessageID: 107
type MspCompGps struct {
	// DistanceToHome is the distance to home in meters.
	DistanceToHome uint16
	// DirectionToHome is the direction to home in degrees from -180 to 180.
	DirectionToHome uint16
	// GpsUpdate is a flag to indicate when a new gps frame is received.
	GpsUpdate uint8
}

// New builds a new instance of the packet.
func (p *MspCompGps) New() Packet {
	return &MspCompGps{}
}

// GetID returns this packet ID.
func (p *MspCompGps) GetID() uint8 {
	return 107
}

// Marshal marshals the packet to a byte array.
func (p *MspCompGps) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.DistanceToHome,
		&p.DirectionToHome,
		&p.GpsUpdate,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspCompGps) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.DistanceToHome,
		&p.DirectionToHome,
		&p.GpsUpdate,
	)
}

func init() {
	registerPacketType(&MspCompGps{})
}
