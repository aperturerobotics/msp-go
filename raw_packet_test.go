package msp

import (
	"bytes"
	"testing"
)

// TestPacketReadWrite tests packet serialization.
func TestPacketReadWrite(t *testing.T) {
	data := []byte{0x04, 0x05, 0x06}
	p := NewRawPacket(0x08, false, data)

	var buf bytes.Buffer
	if _, err := p.WriteTo(&buf); err != nil {
		t.Fatal(err.Error())
	}

	if bytes.Compare(
		buf.Bytes(),
		[]byte{36, 77, 60, 3, 8, 4, 5, 6, 12},
	) != 0 {
		t.Fail()
	}

	pkt, err := ReadRawPacket(&buf)
	if err != nil {
		t.Fatal(err.Error())
	}

	if bytes.Compare(pkt.getDirectionID(), dirToMWC) != 0 {
		t.Fail()
	}

	if bytes.Compare(pkt.GetData(), data) != 0 {
		t.Fail()
	}
}
