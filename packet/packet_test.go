package packet

import (
	"testing"
)

func TestSerializePacket(t *testing.T) {
	for msgID, msgPkt := range packetTypes {
		if msgID != msgPkt.GetID() {
			t.Fail()
		}
	}

	msg := &MspSetMisc{MinThrottle: 100}
	msgSerialized, err := msg.Marshal()
	if err != nil {
		t.Fatal(err.Error())
	}

	nmsg := msg.New()
	if err := nmsg.Unmarshal(msgSerialized); err != nil {
		t.Fatal(err.Error())
	}

	if nmsg.(*MspSetMisc).MinThrottle != 100 {
		t.Fail()
	}
}
