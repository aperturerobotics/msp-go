package packet

import "github.com/pkg/errors"

var errNotImplemented = errors.New("packet not implemented")

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

// ToRaw serializes a packet to a raw packet.
func ToRaw(packet Packet, isRecv, isUnrecognized bool) (*RawPacket, error) {
	data, err := packet.Marshal()
	if err != nil {
		return nil, err
	}

	return NewRawPacket(packet.GetID(), isRecv, isUnrecognized, data), nil
}

// FromRaw decodes a raw packet to a packet.
func FromRaw(rawPacket *RawPacket) (Packet, error) {
	pid := rawPacket.GetID()
	pt, ok := packetTypes[pid]
	if !ok {
		return nil, errors.Errorf("unknown packet type: %d", pid)
	}

	pak := pt.New()
	if err := pak.Unmarshal(rawPacket.GetData()); err != nil {
		return nil, err
	}

	return pak, nil
}
