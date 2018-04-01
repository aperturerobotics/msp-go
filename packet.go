package msp

// Packet contains a defined packet for the MSP protocol.
type Packet interface {
	// New builds a new instance of the packet.
	New() Packet
	// GetID returns this packet ID.
	GetID() uint8
	// Marshal marshals the packet to a byte array.
	Marshal() ([]byte, error)
	// Unmarshal parses the byte array, filling the packet values.
	Unmarshal([]byte) error
}

// packetTypes contains the known packet types.
var packetTypes = map[uint8]Packet{}

// registerPacketType registers a packet type.
func registerPacketType(packet Packet) {
	packetTypes[packet.GetID()] = packet
}

// PacketToRaw serializes a packet to a raw packet.
func PacketToRaw(packet Packet, isRecv bool) (*RawPacket, error) {
	data, err := packet.Marshal()
	if err != nil {
		return nil, err
	}

	return NewRawPacket(packet.GetID(), isRecv, data), nil
}
