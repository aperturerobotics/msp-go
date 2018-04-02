package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aperturerobotics/msp-go"
	"github.com/aperturerobotics/msp-go/serial"

	ts "github.com/tarm/serial"
	"github.com/urfave/cli"
)

var cliCommands []cli.Command
var cliFlags []cli.Flag

func main() {
	app := cli.NewApp()
	app.Name = "msp"
	app.Usage = "msp debug cli / examples"
	app.Commands = cliCommands
	app.Flags = cliFlags
	app.HideVersion = true

	app.Before = openDevice
	app.After = closeDevice

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

var device *msp.Device
var devicePath string
var deviceBaud int
var serialPort *ts.Port

func init() {
	cliFlags = append(
		cliFlags,
		cli.StringFlag{
			Name:        "device-path",
			Usage:       "serial device path",
			Value:       "/dev/ttyUSB0",
			Destination: &devicePath,
		},
		cli.IntFlag{
			Name:        "device-baud",
			Usage:       "serial device baudrate",
			Value:       115200,
			Destination: &deviceBaud,
		},
	)
}

// openDevice opens the MSP device.
func openDevice(ctx *cli.Context) (err error) {
	device, serialPort, err = serial.OpenDeviceSerial(
		context.Background(),
		devicePath,
		deviceBaud,
	)

	return
}

// closeDevice closes the serial device.
func closeDevice(ctx *cli.Context) error {
	if serialPort != nil {
		serialPort.Close()
	}

	return nil
}

func dumpJSON(item interface{}) {
	dat, _ := json.Marshal(item)
	fmt.Printf("%s\n", string(dat))
}
