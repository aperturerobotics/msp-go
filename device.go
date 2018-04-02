package msp

import (
	"context"
	"io"
	"sync"

	"github.com/aperturerobotics/msp-go/packet"
)

// Device represents a remote MSP device.
type Device struct {
	tpt      io.ReadWriter
	writeMtx sync.Mutex
	readMtx  sync.Mutex
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

// WritePacket sends a packet to the device.
func (d *Device) WritePacket(pkt packet.Packet) error {
	rawPkt, err := packet.ToRaw(pkt, false)
	if err != nil {
		return err
	}

	d.writeMtx.Lock()
	_, err = rawPkt.WriteTo(d.tpt)
	d.writeMtx.Unlock()

	return err
}

// ReadPacket reads a packet from the device.
func (d *Device) ReadPacket() (packet.Packet, error) {
	d.readMtx.Lock()
	rawPkt, err := packet.ReadRawPacket(d.tpt)
	d.readMtx.Unlock()
	if err != nil {
		return nil, err
	}

	return packet.FromRaw(rawPkt)
}
