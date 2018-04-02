package main

import (
	"context"
	"fmt"

	"github.com/aperturerobotics/msp-go/packet"
	"github.com/urfave/cli"
)

var sendRCArgs struct {
	channels [packet.ChannelCount]int
}

func init() {
	var channelFlags []cli.Flag
	packet.ForEachChannel(func(c packet.ChannelID) bool {
		cStr := c.String()
		channelFlags = append(channelFlags, cli.IntFlag{
			Name:        cStr,
			Usage:       fmt.Sprintf("sets the rc %s channel (%d) value", cStr, int(c)),
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

func sendRCCommands(ctx *cli.Context) error {
	pkt := &packet.MspRcSet{}
	for chi, chv := range sendRCArgs.channels {
		pkt.SetChannel(packet.ChannelID(chi), uint16(chv))
	}

	return device.Request(context.Background(), pkt)
}
