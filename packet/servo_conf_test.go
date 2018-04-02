package packet

import (
	"testing"
)

func TestMspServoConf(t *testing.T) {
	conf := &MspServoConf{}
	conf.ServoConfs[1].Middle = 10

	dat, err := conf.Marshal()
	if err != nil {
		t.Fatal(err.Error())
	}

	pktb := conf.New().(*MspServoConf)
	if err := pktb.Unmarshal(dat); err != nil {
		t.Fatal(err.Error())
	}

	if pktb.ServoConfs[1].Middle != 10 {
		t.Fail()
	}
}
