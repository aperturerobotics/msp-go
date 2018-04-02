package packet

import "bytes"

// MspIdent contains autopilot identification data.
// Direction: FC->
// MessageID: 100
type MspIdent struct {
	// Version is the version of MultiWii
	Version uint8
	// MultiType is the type of MultiWii
	MultiType uint8
	// MspVersion is the MSP version, not used currently.
	MspVersion uint8
	// Capability is a 32 bit variable to indicate capability of the board.
	Capability uint32
}

// New builds a new instance of the packet.
func (p *MspIdent) New() Packet {
	return &MspIdent{}
}

// GetID returns this packet ID.
func (p *MspIdent) GetID() uint8 {
	return 100
}

// Marshal marshals the packet to a byte array.
func (p *MspIdent) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.Version,
		&p.MultiType,
		&p.MspVersion,
		&p.Capability,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspIdent) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.Version,
		&p.MultiType,
		&p.MspVersion,
		&p.Capability,
	)
}

func init() {
	registerPacketType(&MspIdent{})
}
