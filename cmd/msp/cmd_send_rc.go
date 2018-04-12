package main

import (
	"context"
	"fmt"
	// 	"time"

	"github.com/aperturerobotics/msp-go/packet"
	"github.com/urfave/cli"
)

var sendRCArgs struct {
	channels [packet.ChannelCount]int
}

func init() {
	for i := range sendRCArgs.channels {
		sendRCArgs.channels[i] = 1500
	}

	var channelFlags []cli.Flag
	packet.ForEachChannel(func(c packet.ChannelID) bool {
		cStr := c.String()
		channelFlags = append(channelFlags, cli.IntFlag{
			Name:        cStr,
			Usage:       fmt.Sprintf("sets the rc %s channel (%d) value", cStr, int(c)),
			Value:       sendRCArgs.channels[c],
			Destination: &sendRCArgs.channels[c],
		})
		return true
	})

	cliCommands = append(cliCommands, cli.Command{
		Name:   "send-rc",
		Usage:  "send rc commands over msp",
		Action: sendRCCommands,
		Flags:  channelFlags,
	})
}

func sendRCCommands(cctx *cli.Context) error {
	ctx := context.Background()
	pkt := &packet.MspRcSet{}

	for chi, chv := range sendRCArgs.channels {
		pkt.SetChannel(packet.ChannelID(chi), uint16(chv))
	}

	// timer := time.NewTicker(time.Millisecond * 300)
	// for {
	// <-timer.C

	if err := device.Request(ctx, pkt); err != nil {
		return err
	}
	// }

	/*
		pktrc := &packet.MspRc{}
		return reqDumpPacket(pktrc)
	*/
	return nil
}
