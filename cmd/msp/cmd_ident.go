package main

import (
	"context"

	"github.com/aperturerobotics/msp-go/packet"
	"github.com/urfave/cli"
)

func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name:   "ident",
		Usage:  "identify the connected device",
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
	var pkt packet.Packet = &packet.MspStatus{}
	if err := reqDumpPacket(pkt); err != nil {
		return err
	}

	pkt = &packet.MspPidNames{}
	if err := reqDumpPacket(pkt); err != nil {
		return err
	}

	return nil
}
