package msp

import (
	"context"
	"io"
)

// Device represents a remote MSP device.
type Device struct {
	tpt io.ReadWriter
}

// OpenDevice attempts to open and associate with a device over a read/writer.
func OpenDevice(ctx context.Context, conn io.ReadWriter) (*Device, error) {
	dev := &Device{tpt: conn}
	if err := dev.initConnection(ctx); err != nil {
		return nil, err
	}

	return dev, nil
}

// initConnection attempts to identify the MSP device.
func (d *Device) initConnection(ctx context.Context) error {
	return nil
}
