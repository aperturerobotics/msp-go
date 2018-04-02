package packet

import (
	"testing"
)

// TestMspBoxNames tests the box names packet.
func TestMspBoxNames(t *testing.T) {
	pkt := &MspBoxNames{BoxNames: []string{"test1", "test2", "test3"}}
	dat, err := pkt.Marshal()
	if err != nil {
		t.Fatal(err.Error())
	}

	pktb := pkt.New()
	if err := pktb.Unmarshal(dat); err != nil {
		t.Fatal(err.Error())
	}

	pktbn := pktb.(*MspBoxNames)
	if len(pktbn.BoxNames) != 3 || pktbn.BoxNames[1] != pkt.BoxNames[1] {
		t.Fail()
	}
}
