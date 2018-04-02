package packet

import (
	"bytes"
)

// MspSetWp contains waypoint data for the FC.
// Direction: ->FC
// MessageID: 209
type MspSetWp struct {
	// SetWpNo is the number of waypoints.
	SetWpNo uint8
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
func (p *MspSetWp) New() Packet {
	return &MspSetWp{}
}

// GetID returns this packet ID.
func (p *MspSetWp) GetID() uint8 {
	return 209
}

// Marshal marshals the packet to a byte array.
func (p *MspSetWp) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.SetWpNo,
		&p.Lat,
		&p.Lng,
		&p.AltHold,
		&p.Heading,
		&p.TimeToStay,
		&p.NavFlag,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetWp) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.SetWpNo,
		&p.Lat,
		&p.Lng,
		&p.AltHold,
		&p.Heading,
		&p.TimeToStay,
		&p.NavFlag,
	)
}

func init() {
	registerPacketType(&MspSetWp{})
}
