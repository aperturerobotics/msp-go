package packet

import (
	"bytes"
)

// MspBoxSet contains BOX data to the FC.
// The size of the message is enough to know the number of BOX.
// For each BOX, there is a 16 bit variable which indicates the AUX1->AUX4 activation switch.
// Bit 1: AUX1 LOW state / bit 2: AUX1 MID state / bit 3: AUX1 HIGH state / bit 4: AUX2 LOW state ... bit 13: AUX 4 HIGH state
// Direction: ->FC
// MessageID: 203
type MspBoxSet struct {
	// BoxSetItems contains the box items.
	BoxSetItems []uint16
}

// New builds a new instance of the packet.
func (p *MspBoxSet) New() Packet {
	return &MspBoxSet{}
}

// GetID returns this packet ID.
func (p *MspBoxSet) GetID() uint8 {
	return 203
}

// Marshal marshals the packet to a byte array.
func (p *MspBoxSet) Marshal() ([]byte, error) {
	datas := make([]interface{}, len(p.BoxSetItems))
	for i, v := range p.BoxSetItems {
		datas[i] = v
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspBoxSet) Unmarshal(data []byte) error {
	p.BoxSetItems = make([]uint16, len(data)/2)
	datas := make([]interface{}, len(p.BoxSetItems))
	for i := range datas {
		datas[i] = &p.BoxSetItems[i]
	}

	return readLSB(
		bytes.NewReader(data),

		datas...,
	)
}

func init() {
	registerPacketType(&MspBoxSet{})
}
