package main

import (
	"bytes"
	"encoding/binary"
	"net"
	"testing"
)

// TestReadFrame checks that a frame can be loaded from the network
func TestReadFrame(t *testing.T) {
	side1, side2 := net.Pipe()
	defer side1.Close()
	defer side2.Close()

	// NOTE: Yes, I know this might lead to a race condition but since it only
	// actually is looked at after we read a frame it should™ be safe.
	var sender bytes.Buffer
	go func() {
		sender.Write([]byte("some random text"))

		var frameLength [4]byte
		binary.LittleEndian.PutUint32(frameLength[:], uint32(sender.Len()))

		var buffer bytes.Buffer
		buffer.Write([]byte{MagicByte})
		buffer.Write(frameLength[:])
		buffer.Write(sender.Bytes())

		// Write to the first side
		side1.Write(buffer.Bytes())
	}()

	buf := make([]byte, 1024)
	frame, err := readFrame(side2, buf)
	if err != nil {
		t.Error(err)
		return
	}

	if frame.Magic != MagicByte {
		t.Errorf("Failed to verify magic byte. Got: %v. Expected: %v", frame.Magic, MagicByte)
	}

	if frame.Length != uint32(sender.Len()) {
		t.Errorf("Failed to verify len. Got: %v. Expected: %v", frame.Length, sender.Len())
	}

	if "some random text" != string(frame.Payload) {
		t.Errorf("Failed to verify payload. Got: %s. Expected: %s", string(frame.Payload), "some random text")
	}
}
