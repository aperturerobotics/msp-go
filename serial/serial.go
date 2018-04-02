package serial

import (
	"context"

	"github.com/aperturerobotics/msp-go"
	ts "github.com/tarm/serial"
)

// OpenDeviceSerial opens a device on a serial port.
func OpenDeviceSerial(
	ctx context.Context,
	port string,
	baudRate int,
) (*msp.Device, *ts.Port, error) {
	conf := &ts.Config{Name: port, Baud: baudRate}

	s, err := ts.OpenPort(conf)
	if err != nil {
		return nil, nil, err
	}

	dev, err := msp.OpenDevice(ctx, s)
	if err != nil {
		_ = s.Close()
		return nil, nil, err
	}

	return dev, s, nil
}
