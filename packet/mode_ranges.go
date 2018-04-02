package packet

import (
	"bytes"
)

// MspMode is a mode in the MSP system.
type MspMode struct {
	// PermanentID is the ID of the mode.
	PermanentID uint8
	// AuxChannelIndex is the aux channel index of the mode.
	AuxChannelIndex uint8
	// RangeStartStep is the start value for this element in blocks of 25 where 0 -> 900 and 48 -> 2100.
	RangeStartStep uint8
	// RangeEndStep is the end value of this element in blocks of 25 where 0 -> 900 and 48 -> 2100.
	RangeEndStep uint8
}

// MspModeRanges contains the current auxiliary mode settings from the FC.
// Direction: FC->
// MessageID: 34
type MspModeRanges struct {
	// Modes is the list of modes.
	Modes []MspMode
}

// New builds a new instance of the packet.
func (p *MspModeRanges) New() Packet {
	return &MspModeRanges{}
}

// GetID returns this packet ID.
func (p *MspModeRanges) GetID() uint8 {
	return 34
}

func (p *MspModeRanges) getValInterfaces() []interface{} {
	vals := make([]interface{}, len(p.Modes)*4)
	for ix, mode := range p.Modes {
		im := ix * 4
		vals[im] = &mode.PermanentID
		vals[im+1] = &mode.AuxChannelIndex
		vals[im+2] = &mode.RangeStartStep
		vals[im+3] = &mode.RangeEndStep
	}
	return vals
}

// Marshal marshals the packet to a byte array.
func (p *MspModeRanges) Marshal() ([]byte, error) {
	return writeLSBBuf(
		p.getValInterfaces()...,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspModeRanges) Unmarshal(data []byte) error {
	p.Modes = make([]MspMode, len(data)/4)

	return readLSB(
		bytes.NewReader(data),

		p.getValInterfaces()...,
	)
}

func init() {
	registerPacketType(&MspModeRanges{})
}
