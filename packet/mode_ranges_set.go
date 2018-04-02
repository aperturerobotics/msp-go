package packet

import (
	"bytes"
)

// MspModeSeq is a msp mode with a sequence number.
type MspModeSeq struct {
	MspMode

	// SequenceID is the sequence ID of the mode.
	SequenceID uint8
}

// MspSetModeRanges sets the current auxiliary mode settings to the FC.
// Direction: ->FC
// MessageID: 35
type MspSetModeRanges struct {
	// Modes is the list of modes.
	Modes []MspModeSeq
}

// New builds a new instance of the packet.
func (p *MspSetModeRanges) New() Packet {
	return &MspSetModeRanges{}
}

// GetID returns this packet ID.
func (p *MspSetModeRanges) GetID() uint8 {
	return 35
}

// InitFromMspModeRanges initializes the values from the existing.
func (p *MspSetModeRanges) InitFromMspModeRanges(ranges *MspModeRanges) {
	p.Modes = make([]MspModeSeq, len(ranges.Modes))
	for i, m := range ranges.Modes {
		p.Modes[i] = MspModeSeq{MspMode: m, SequenceID: uint8(i)}
	}
}

func (p *MspSetModeRanges) getValInterfaces() []interface{} {
	vals := make([]interface{}, len(p.Modes)*5)
	for ix, mode := range p.Modes {
		im := ix * 5
		vals[im] = &mode.SequenceID
		vals[im+1] = &mode.PermanentID
		vals[im+2] = &mode.AuxChannelIndex
		vals[im+3] = &mode.RangeStartStep
		vals[im+4] = &mode.RangeEndStep
	}
	return vals
}

// Marshal marshals the packet to a byte array.
func (p *MspSetModeRanges) Marshal() ([]byte, error) {
	return writeLSBBuf(
		p.getValInterfaces()...,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetModeRanges) Unmarshal(data []byte) error {
	p.Modes = make([]MspModeSeq, len(data)/5)

	return readLSB(
		bytes.NewReader(data),

		p.getValInterfaces()...,
	)
}

func init() {
	registerPacketType(&MspSetModeRanges{})
}
