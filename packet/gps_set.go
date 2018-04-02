package packet

import (
	"bytes"
)

// MspSetRawGps contains GPS data to the FC.
// Direction: ->FC
// MessageID: 201
type MspSetRawGps struct {
	// GpsFix indicates if we have a GPS fix.
	GpsFix uint8
	// NumSat is the number of satellites.
	NumSat uint8
	// Lat is the latitude, 1/10000000 deg
	Lat uint32
	// Lng is the longitude, 1/10000000 deg
	Lng uint32
	// Altitude is the altitude in meters.
	Altitude uint16
	// Speed is the speed in cm/s
	Speed uint16
}

// New builds a new instance of the packet.
func (p *MspSetRawGps) New() Packet {
	return &MspSetRawGps{}
}

// GetID returns this packet ID.
func (p *MspSetRawGps) GetID() uint8 {
	return 201
}

// Marshal marshals the packet to a byte array.
func (p *MspSetRawGps) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.GpsFix,
		&p.NumSat,
		&p.Lat,
		&p.Lng,
		&p.Altitude,
		&p.Speed,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetRawGps) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.GpsFix,
		&p.NumSat,
		&p.Lat,
		&p.Lng,
		&p.Altitude,
		&p.Speed,
	)
}

func init() {
	registerPacketType(&MspSetRawGps{})
}
