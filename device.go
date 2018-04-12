package msp

import (
	"context"
	"io"
	"sync"

	"github.com/aperturerobotics/msp-go/packet"
	"github.com/pkg/errors"
)

// Device represents a remote MSP device.
type Device struct {
	tpt io.ReadWriter

	reqMtx   sync.Mutex
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
	rawPkt, err := packet.ToRaw(pkt, false, false)
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

// Request requests that the MSP fill the packet with data.
func (d *Device) Request(ctx context.Context, pkt packet.Packet) error {
	d.reqMtx.Lock()
	defer d.reqMtx.Unlock()

	if err := d.WritePacket(pkt); err != nil {
		return err
	}

	d.readMtx.Lock()
	rawPkt, err := packet.ReadRawPacket(d.tpt)
	d.readMtx.Unlock()
	if err != nil {
		return err
	}

	if rawPkt.GetID() != pkt.GetID() {
		return errors.Errorf(
			"received unexpected response to packet %d: id %d len %d",
			pkt.GetID(),
			rawPkt.GetID(),
			len(rawPkt.GetData()),
		)
	}

	if rawPkt.GetIsUnrecognized() {
		return errors.Errorf(
			"received unrecognized response to packet %d",
			pkt.GetID(),
		)
	}

	if len(rawPkt.GetData()) > 0 {
		if err := pkt.Unmarshal(rawPkt.GetData()); err != nil {
			return errors.Wrap(err, "unmarshal packet")
		}
	}

	return nil
}
