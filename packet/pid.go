package packet

import (
	"bytes"
)

// PidItem is a pid item.
type PidItem [3]uint8

// MspPid contains PID data from the FC.
// ROLL / PITCH / YAW / ALT / POS / POSR / NAVR / LEVEL /MAG / VEL
// Direction: FC->
// MessageID: 112
type MspPid struct {
	// PidItems contains the pid items.
	PidItems [PidItems]PidItem
}

// New builds a new instance of the packet.
func (p *MspPid) New() Packet {
	return &MspPid{}
}

// GetID returns this packet ID.
func (p *MspPid) GetID() uint8 {
	return 112
}

// Marshal marshals the packet to a byte array.
func (p *MspPid) Marshal() ([]byte, error) {
	datas := make([]interface{}, PidItems*3)
	for i := range p.PidItems {
		for x, v := range p.PidItems[i] {
			datas[(i*3)+x] = v
		}
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspPid) Unmarshal(data []byte) error {
	datas := make([]interface{}, PidItems*3)
	for i := range p.PidItems {
		for x := range p.PidItems[i] {
			datas[(i*3)+x] = &p.PidItems[i][x]
		}
	}

	return readLSB(
		bytes.NewReader(data),

		datas...,
	)
}

func init() {
	registerPacketType(&MspPid{})
}
