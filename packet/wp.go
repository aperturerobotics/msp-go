package packet

import (
	"bytes"
)

// MspWp contains waypoint data from the FC.
// Direction: FC->
// MessageID: 118
type MspWp struct {
	// WpNo is the number of waypoints.
	WpNo uint8
	// Lat is the latitude.
	Lat uint32
	// Lng is the longitude.
	Lng uint32
	// AltHold is the altitude to hold.
	AltHold uint32
	// Heading is the heading to hold.
	Heading uint16
	// TimeToStay is how long to stay at the waypoint.
	TimeToStay uint16
	// NavFlag is any navigation flags.
	NavFlag uint8
}

// New builds a new instance of the packet.
func (p *MspWp) New() Packet {
	return &MspWp{}
}

// GetID returns this packet ID.
func (p *MspWp) GetID() uint8 {
	return 118
}

// Marshal marshals the packet to a byte array.
func (p *MspWp) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.WpNo,
		&p.Lat,
		&p.Lng,
		&p.AltHold,
		&p.Heading,
		&p.TimeToStay,
		&p.NavFlag,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspWp) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.WpNo,
		&p.Lat,
		&p.Lng,
		&p.AltHold,
		&p.Heading,
		&p.TimeToStay,
		&p.NavFlag,
	)
}

func init() {
	registerPacketType(&MspWp{})
}
