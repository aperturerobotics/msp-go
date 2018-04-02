package packet

import (
	"bytes"
)

// MspAnalog contains GPS data from the FC.
// Direction: FC->
// MessageID: 110
type MspAnalog struct {
	// VBat is the voltage of the battery in 1/10 of volt.
	VBat uint8
	// PowerMeterSum is the power meter sum.
	PowerMeterSum uint16
	// Rssi is the rssi of the signal, between 0 and 1023 volts.
	Rssi uint16
	// Amperage is current amperage draw
	Amperage uint16
}

// New builds a new instance of the packet.
func (p *MspAnalog) New() Packet {
	return &MspAnalog{}
}

// GetID returns this packet ID.
func (p *MspAnalog) GetID() uint8 {
	return 110
}

// Marshal marshals the packet to a byte array.
func (p *MspAnalog) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.VBat,
		&p.PowerMeterSum,
		&p.Rssi,
		&p.Amperage,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspAnalog) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.VBat,
		&p.PowerMeterSum,
		&p.Rssi,
		&p.Amperage,
	)
}

func init() {
	registerPacketType(&MspAnalog{})
}
