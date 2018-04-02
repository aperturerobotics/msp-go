package packet

import "fmt"

// ServoCount is the number of servos the FCU talks about.
const ServoCount = 8

// MotorCount is the number of motors the FCU talks about.
const MotorCount = ServoCount

// ChannelCount is the number of RC channels the FCU talks about.
const ChannelCount = ServoCount

// PidItems is the constant number of PID parameter items.
const PidItems = 10

// MotorPinCount is the number of motor pins.
const MotorPinCount = 8

// ChannelID is the identifier / index of the channel.
type ChannelID uint8

const (
	// Roll is the roll channel 1.
	Roll ChannelID = iota
	// Pitch is the pitch channel 2.
	Pitch
	// Yaw is the yaw channel 3.
	Yaw
	// Throttle is the throttle channel 4.
	Throttle
	// Aux1 is the aux channel 5.
	Aux1
	// Aux2 is the second aux channel 6.
	Aux2
	// Aux3 is the third aux channel 7.
	Aux3
	// Aux4 is the fourth aux channel 8.
	Aux4
)

var channelIDFlag = map[ChannelID]string{
	Roll:     "roll",
	Pitch:    "pitch",
	Yaw:      "yaw",
	Throttle: "throttle",
	Aux1:     "aux1",
	Aux2:     "aux2",
	Aux3:     "aux3",
	Aux4:     "aux4",
}

// String returns the string representation of the channel.
func (c ChannelID) String() string {
	val, ok := channelIDFlag[c]
	if !ok {
		return fmt.Sprintf("unknown: %d", c)
	}

	return val
}

// ForEachChannel iterates over the channel ids.
func ForEachChannel(cb func(c ChannelID) bool) {
	for k := range channelIDFlag {
		if !cb(k) {
			return
		}
	}
}
