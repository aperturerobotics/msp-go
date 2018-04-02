package packet

import (
	"bytes"
)

// MspBox contains BOX data from the FC.
// The size of the message is enough to know the number of BOX.
// For each BOX, there is a 16 bit variable which indicates the AUX1->AUX4 activation switch.
// Bit 1: AUX1 LOW state / bit 2: AUX1 MID state / bit 3: AUX1 HIGH state / bit 4: AUX2 LOW state ... bit 13: AUX 4 HIGH state
// Direction: FC->
// MessageID: 113
type MspBox struct {
	// BoxItems contains the box items.
	BoxItems []uint16
}

// New builds a new instance of the packet.
func (p *MspBox) New() Packet {
	return &MspBox{}
}

// GetID returns this packet ID.
func (p *MspBox) GetID() uint8 {
	return 113
}

// Marshal marshals the packet to a byte array.
func (p *MspBox) Marshal() ([]byte, error) {
	datas := make([]interface{}, len(p.BoxItems))
	for i, v := range p.BoxItems {
		datas[i] = v
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspBox) Unmarshal(data []byte) error {
	p.BoxItems = make([]uint16, len(data)/2)
	datas := make([]interface{}, len(p.BoxItems))
	for i := range datas {
		datas[i] = &p.BoxItems[i]
	}

	return readLSB(
		bytes.NewReader(data),

		datas...,
	)
}

func init() {
	registerPacketType(&MspBox{})
}
