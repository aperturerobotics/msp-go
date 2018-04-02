package packet

import (
	"bytes"
)

// MspMisc contains misc data from the FC.
// Direction: FC->
// MessageID: 114
type MspMisc struct {
	// PowerTrigger is the power trigger value.
	PowerTrigger uint16
	// MinThrottle is the minimum throttle to run in idle state, 1000-2000.
	MinThrottle uint16
	// MaxThrottle is the maximum throttle in 1000-2000.
	MaxThrottle uint16
	// MinCommand is the throttle at the lowest position, 1000-2000.
	MinCommand uint16
	// FailsafeThrottle is the throttle to use in failsafe, 1000-2000.
	FailsafeThrottle uint16
	// ArmCounter is the number of times the system has been armed.
	ArmCounter uint16
	// Lifetime is the amount of time the system has been online.
	Lifetime uint32
	// MagDeclination is the magnetic declination in 1/10 degree.
	MagDeclination uint16
	// VBatScale is the battery voltage scale.
	VBatScale uint8
	// VBatWarn1 is the first warning voltage in 1/10 volts.
	VBatWarn1 uint8
	// VBatWarn2 is the first warning voltage in 1/10 volts.
	VBatWarn2 uint8
	// VBatCrit is the critical warning voltage in 1/10 volts.
	VBatCrit uint8
}

// New builds a new instance of the packet.
func (p *MspMisc) New() Packet {
	return &MspMisc{}
}

// GetID returns this packet ID.
func (p *MspMisc) GetID() uint8 {
	return 114
}

// Marshal marshals the packet to a byte array.
func (p *MspMisc) Marshal() ([]byte, error) {
	return writeLSBBuf(
		&p.PowerTrigger,
		&p.MinThrottle,
		&p.MaxThrottle,
		&p.MinCommand,
		&p.FailsafeThrottle,
		&p.ArmCounter,
		&p.Lifetime,
		&p.MagDeclination,
		&p.VBatScale,
		&p.VBatWarn1,
		&p.VBatWarn2,
		&p.VBatCrit,
	)
}

// Unmarshal parses the byte array, filling the packet values.
func (p *MspMisc) Unmarshal(data []byte) error {
	return readLSB(
		bytes.NewReader(data),

		&p.PowerTrigger,
		&p.MinThrottle,
		&p.MaxThrottle,
		&p.MinCommand,
		&p.FailsafeThrottle,
		&p.ArmCounter,
		&p.Lifetime,
		&p.MagDeclination,
		&p.VBatScale,
		&p.VBatWarn1,
		&p.VBatWarn2,
		&p.VBatCrit,
	)
}

func init() {
	registerPacketType(&MspMisc{})
}
