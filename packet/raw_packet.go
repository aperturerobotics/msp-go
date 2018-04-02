package packet

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
	// dirUnrecognized is used when the message is unknown
	dirUnrecognized = []byte(`!`)
)

// RawPacket is a MSP packet ready to write to a stream.
type RawPacket struct {
	packetID     uint8
	recv         bool
	unrecognized bool
	data         []byte
}

// NewRawPacket builds a new raw packet.
func NewRawPacket(packetID uint8, isRecv bool, isUnrecognized bool, data []byte) *RawPacket {
	return &RawPacket{
		packetID:     packetID,
		recv:         isRecv,
		data:         data,
		unrecognized: isUnrecognized,
	}
}

// GetID returns the packet ID.
func (p *RawPacket) GetID() uint8 {
	return p.packetID
}

// GetIsRecv indicates if this is a packet received from the FC.
func (p *RawPacket) GetIsRecv() bool {
	return p.recv
}

// GetIsUnrecognized indicates if this packet indicated the peer did not recognize the packet.
func (p *RawPacket) GetIsUnrecognized() bool {
	return p.unrecognized
}

// GetData returns the in-band data.
func (p *RawPacket) GetData() []byte {
	return p.data
}

// readLSB reads with the correct endian conversion.
func readLSB(r io.Reader, datas ...interface{}) error {
	for _, data := range datas {
		if err := binary.Read(r, binary.LittleEndian, data); err != nil {
			return err
		}
	}

	return nil
}

// writeLSB writes with the correct endian conversion.
func writeLSB(w io.Writer, datas ...interface{}) error {
	for _, data := range datas {
		if err := binary.Write(w, binary.LittleEndian, data); err != nil {
			return err
		}
	}

	return nil
}

// writeLSBBuf writes with the correct endian conversion to a buffep.
func writeLSBBuf(datas ...interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := writeLSB(&buf, datas...); err != nil {
		return nil, err
	}

	return (&buf).Bytes(), nil
}

// ReadRawPacket attempts to read a raw packet from the reader.
func ReadRawPacket(r io.Reader) (rp *RawPacket, err error) {
	defer func() {
		var isRecognized bool
		if rp != nil {
			isRecognized = !rp.GetIsUnrecognized()
		}
		// fmt.Printf("read packet with err: %v is recognized: %v\n", err, isRecognized)
	}()

	// Wait until we see the magic preamble.
	buf := make([]byte, 2)
	for {
		buf[0] = buf[1]
		err := readLSB(r, &buf[1])
		if err != nil {
			return nil, err
		}

		if bytes.Compare(buf, preamble) == 0 {
			break
		}
	}

	// Read the direction ID
	var dirID byte
	if err := readLSB(r, &dirID); err != nil {
		return nil, err
	}

	var isRecv bool
	var isUnrecognized bool
	switch dirID {
	case dirUnrecognized[0]:
		isUnrecognized = true
		fallthrough
	case dirFromMWC[0]:
		isRecv = true
	default:
		return nil, errors.Errorf(
			"unexpected direction indicator: %v (%s)",
			dirID,
			string([]rune{rune(dirID)}),
		)
	case dirToMWC[0]:
	}

	// Read the size and command
	var messageID uint8
	var dataSize uint8
	if err := readLSB(r, &dataSize, &messageID); err != nil {
		return nil, err
	}

	// Cast the size to a uint8
	if dataSize > maxPacketSize {
		return nil, errors.Errorf("unexpected data size: %v", dataSize)
	}

	// Read data
	dataBuf := make([]byte, dataSize)
	if err := readLSB(r, dataBuf); err != nil {
		return nil, err
	}

	// Read CRC
	var crc uint8
	if err := readLSB(r, &crc); err != nil {
		return nil, err
	}

	// Build packet
	pkt := NewRawPacket(messageID, isRecv, isUnrecognized, dataBuf)
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

// WriteTo serializes the packet and writes it to a writep.
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
