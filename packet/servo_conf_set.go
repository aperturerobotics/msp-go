package packet

import (
	"bytes"
)

// MspSetServoConf contains servo configuration data for the FC.
// Direction: ->FC
// MessageID: 212
type MspSetServoConf struct {
	// ServoConfs contains servo configuration data.
	ServoConfs [ServoCount]ServoConf
}

// New builds a new instance of the packet.
func (p *MspSetServoConf) New() Packet {
	return &MspSetServoConf{}
}

// GetID returns this packet ID.
func (p *MspSetServoConf) GetID() uint8 {
	return 212
}

// Marshal marshals the packet to a byte array.
func (p *MspSetServoConf) Marshal() ([]byte, error) {
	datas := make([]interface{}, ServoCount*4)
	for i := range p.ServoConfs {
		copy(datas[i*4:], (&p.ServoConfs[i]).GetItems())
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspSetServoConf) Unmarshal(data []byte) error {
	datas := make([]interface{}, ServoCount*4)
	for i := range p.ServoConfs {
		copy(datas[i*4:], (&p.ServoConfs[i]).GetItems())
	}

	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspSetServoConf{})
}
