package msp

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

var maxPacketSize uint8 = 255

var (
	// preamble is sent before each packet.
	preamble = []byte(`$M`)
	// dirToMWC is used when writing to the MWC
	dirToMWC = []byte(`<`)
	// dirFromMWC is used when reading from the MWC
	dirFromMWC = []byte(`>`)
)

// RawPacket is a MSP packet ready to write to a stream.
type RawPacket struct {
	packetID uint8
	recv     bool
	data     []byte
}

// NewRawPacket builds a new raw packet.
func NewRawPacket(packetID uint8, isRecv bool, data []byte) *RawPacket {
	return &RawPacket{
		packetID: packetID,
		recv:     isRecv,
		data:     data,
	}
}

// readLSB reads with the correct endian conversion.
func readLSB(r io.Reader, data []byte) error {
	return binary.Read(r, binary.LittleEndian, data)
}

// ReadRawPacket attempts to read a raw packet from the reader.
func ReadRawPacket(r io.Reader) (*RawPacket, error) {
	// Wait until we see the magic preamble.
	buf := make([]byte, 2)
	for {
		buf[0] = buf[1]
		err := readLSB(r, buf[1:])
		if err != nil {
			return nil, err
		}

		if bytes.Compare(buf, preamble) == 0 {
			break
		}
	}

	// Read the direction ID
	if err := readLSB(r, buf[0:1]); err != nil {
		return nil, err
	}

	var isRecv bool
	switch buf[0] {
	case dirFromMWC[0]:
		isRecv = true
	default:
		return nil, errors.Errorf("unexpected direction indicator: %v", buf[0])
	case dirToMWC[0]:
	}

	// Read the size
	if err := readLSB(r, buf[0:1]); err != nil {
		return nil, err
	}

	// Cast the size to a uint8
	dataSize := uint8(buf[0])
	if dataSize > maxPacketSize {
		return nil, errors.Errorf("unexpected data size: %v", dataSize)
	}

	// Read the command
	if err := readLSB(r, buf[0:1]); err != nil {
		return nil, err
	}

	messageID := uint8(buf[0])

	// Read data
	dataBuf := make([]byte, dataSize)
	if err := readLSB(r, dataBuf); err != nil {
		return nil, err
	}

	// Read CRC
	if err := readLSB(r, buf[0:1]); err != nil {
		return nil, err
	}

	crc := uint8(buf[0])

	// Build packet
	pkt := NewRawPacket(messageID, isRecv, dataBuf)
	if expectedCrc := pkt.getCrc(); expectedCrc != crc {
		return nil, errors.Errorf(
			"incoming crc %v != expected %v data len %d",
			crc,
			expectedCrc,
			dataSize,
		)
	}

	return pkt, nil
}

// WriteTo serializes the packet and writes it to a writer.
func (p *RawPacket) WriteTo(w io.Writer) (n int64, err error) {
	return p.writeFuncs([]func() (n int, err error){
		// Preamble
		func() (n int, err error) { return w.Write(preamble) },
		func() (n int, err error) { return w.Write(p.getDirectionID()) },
		p.writeLSBFunc(w, uint8(len(p.data))),
		p.writeLSBFunc(w, p.packetID),
		p.writeLSBFunc(w, p.data),
		p.writeLSBFunc(w, p.getCrc()),
	})
}

// GetData returns the inner packet data.
func (p *RawPacket) GetData() []byte {
	return p.data
}

// getDirectionID returns the code for the direction this packet is going.
func (p *RawPacket) getDirectionID() []byte {
	if p.recv {
		return dirFromMWC
	}

	return dirToMWC
}

// getCrc returns the crc bytes for the data.
func (p *RawPacket) getCrc() uint8 {
	crc := uint8(len(p.data)) ^ p.packetID
	for _, d := range p.data {
		crc = crc ^ uint8(d)
	}

	return crc
}

// writeLSBFunc returns a function that writes the data in lsb.
func (p *RawPacket) writeLSBFunc(
	w io.Writer,
	data interface{},
) func() (n int, err error) {
	return func() (n int, err error) {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, data); err != nil {
			return 0, err
		}

		ni, erri := (&buf).WriteTo(w)
		return int(ni), erri
	}
}

// writeFuncs writes a set of data to an output
func (p *RawPacket) writeFuncs(
	fns []func() (n int, err error),
) (n int64, err error) {
	for _, fn := range fns {
		var ni int
		ni, err = fn()
		n += int64(ni)
		if err != nil {
			return
		}
	}

	return
}