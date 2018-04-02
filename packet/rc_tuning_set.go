package packet

import (
	"bytes"
)

// MspSetRcTuning contains RC tuning data to the FC.
// Direction: ->FC
// MessageID: 204
type MspSetRcTuning struct {
	// RcRate is the RC rate.
	RcRate uint8
	// RcExpo is the RC expo.
	RcExpo uint8
	// RollPitchRate is the roll/pitch rate.
	RollPitchRate uint8
	// YawRate is the yaw rate.
	YawRate uint8
	// DynThrPID indicates dynamic throttle PIDs.
	DynThrPID uint8
	// ThrottleMid is the mid value of the throttle.
	ThrottleMid uint8
	// ThrottleExpo is the throttle exponential.
	ThrottleExpo uint8
}

// New builds a new instance of the packet.
func (p *MspSetRcTuning) New() Packet {
	return &MspSetRcTuning{}
}

// GetID returns this packet ID.
func (p *MspSetRcTuning) GetID() uint8 {
	return 204
}

// Marshal marshals the packet to a byte array.
func (p *MspSetRcTuning) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.RcRate,
		&p.RcExpo,
		&p.RollPitchRate,
		&p.YawRate,
		&p.DynThrPID,
		&p.ThrottleMid,
		&p.ThrottleExpo,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetRcTuning) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.RcRate,
		&p.RcExpo,
		&p.RollPitchRate,
		&p.YawRate,
		&p.DynThrPID,
		&p.ThrottleMid,
		&p.ThrottleExpo,
	)
}

func init() {
	registerPacketType(&MspSetRcTuning{})
}
