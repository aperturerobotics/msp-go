package main

import (
	"context"
	"fmt"

	"github.com/aperturerobotics/msp-go/packet"
	"github.com/urfave/cli"
)

func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name:   "ident",
		Usage:  "identify and debug the connected device",
		Action: identifyDevice,
	})
}

func reqDumpPacket(pkt packet.Packet) error {
	if err := device.Request(context.Background(), pkt); err != nil {
		return err
	}

	dumpJSON(pkt)
	return nil
}

func identifyDevice(ctx *cli.Context) error {
	pkts := []packet.Packet{
		&packet.MspStatus{},
		&packet.MspPidNames{},
		&packet.MspModeRanges{},
		&packet.MspAltitude{},
		&packet.MspAttitude{},
		&packet.MspRc{},
		&packet.MspAnalog{},
	}

	for _, pkt := range pkts {
		if err := reqDumpPacket(pkt); err != nil {
			fmt.Printf("packet %#v failed: %v\n", pkt, err.Error())
			// return err
		}
	}

	return nil
}
