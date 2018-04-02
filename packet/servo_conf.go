package packet

import (
	"bytes"
)

// ServoConf contains servo configuration data.
type ServoConf struct {
	// Min is the minimum value of the servo
	// Range is 1000-2000
	Min uint16
	// Max is the maximum value of the servo
	// Range is 1000-2000
	Max uint16
	// Middle is the neutral value of the servo
	// Range is 1000-2000
	Middle uint16
	// Rate is the servo rate to use.
	// Between 0-100.
	Rate uint8
}

// GetItems returns the items to encode/decode
func (c *ServoConf) GetItems() []interface{} {
	return []interface{}{
		&c.Min,
		&c.Max,
		&c.Middle,
		&c.Rate,
	}
}

// MspServoConf contains servo configuration data.
// Direction: FC->
// MessageID: 120
type MspServoConf struct {
	// ServoConfs contains servo configuration data.
	ServoConfs [ServoCount]ServoConf
}

// New builds a new instance of the packet.
func (p *MspServoConf) New() Packet {
	return &MspServoConf{}
}

// GetID returns this packet ID.
func (p *MspServoConf) GetID() uint8 {
	return 120
}

// Marshal marshals the packet to a byte array.
func (p *MspServoConf) Marshal() ([]byte, error) {
	datas := make([]interface{}, ServoCount*4)
	for i := range p.ServoConfs {
		copy(datas[i*4:], (&p.ServoConfs[i]).GetItems())
	}

	return writeLSBBuf(datas...)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspServoConf) Unmarshal(data []byte) error {
	datas := make([]interface{}, ServoCount*4)
	for i := range p.ServoConfs {
		copy(datas[i*4:], (&p.ServoConfs[i]).GetItems())
	}

	return readLSB(bytes.NewReader(data), datas...)
}

func init() {
	registerPacketType(&MspServoConf{})
}
