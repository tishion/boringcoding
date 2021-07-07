package lightrpc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const (
	SIGNATURE_LEN = 4
	VERSION_LEN   = 2
	LENGTH_LEN    = 4
	HEADER_LEN    = SIGNATURE_LEN + VERSION_LEN + LENGTH_LEN

	MAX_DATA_LEN = 0xfffe

	SIG_VALUE = 0x4C525043 // 'LRPC'
	VER_VALUE = 0x0001     // 0.1
)

type packetHeader struct {
	Signature uint32 // Message signature
	Version   uint16 // Message version
	Length    uint32 // Message Data Length
}

func packetHeaderSplitter(data []byte, eof bool) (advance int, token []byte, err error) {
	if !eof && len(data) == HEADER_LEN {
		var signature uint32 = 0
		var version uint16 = 0
		var dataLen uint32 = 0
		var offset int = 0
		// read and validate the signature
		binary.Read(bytes.NewBuffer(data[offset:offset+SIGNATURE_LEN]), binary.BigEndian, &signature)
		if signature != SIG_VALUE {
			return
		}
		offset += SIGNATURE_LEN

		// read and validate the version
		binary.Read(bytes.NewBuffer(data[offset:offset+VERSION_LEN]), binary.BigEndian, &version)
		if version != VER_VALUE {
			return
		}
		offset += VERSION_LEN

		// read and validate the data length
		binary.Read(bytes.NewBuffer(data[offset:offset+LENGTH_LEN]), binary.BigEndian, &dataLen)
		if dataLen > MAX_DATA_LEN {
			return
		}

		// extact the packet header only
		return HEADER_LEN, data[:HEADER_LEN], nil
	}
	return
}

type Conn struct {
	io io.ReadWriter
}

func LightRpcConn(rw io.ReadWriter) Conn {
	return Conn{
		io: rw,
	}
}

func (c Conn) Send(data []byte) error {
	if len(data) > MAX_DATA_LEN {
		return errors.New("data exceeds the length limit of the rpc packet")
	}

	return c.sendPacketWithData(data)
}

func (c Conn) Receive() ([]byte, error) {
	// create the scanner
	s := bufio.NewScanner(c.io)

	// set scanner buffer
	b := make([]byte, HEADER_LEN)
	s.Buffer(b, HEADER_LEN)

	// set splitter for the scanner to extract the packet header
	s.Split(packetHeaderSplitter)

	ok := s.Scan()
	if !ok {
		return nil, s.Err()
	}

	h := packetHeader{
		Signature: 0,
		Version:   0,
		Length:    0,
	}
	r := bytes.NewReader(s.Bytes())

	// packetHeader read done, deserialize it
	var err error
	err = binary.Read(r, binary.BigEndian, &h.Signature)
	if nil != err {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &h.Version)
	if nil != err {
		return nil, err
	}

	err = binary.Read(r, binary.BigEndian, &h.Length)
	if nil != err {
		return nil, err
	}

	d := make([]byte, h.Length)
	c.io.Read(d)

	return d, nil
}

func (c Conn) sendPacketWithData(data []byte) error {
	h := packetHeader{
		Signature: SIG_VALUE,
		Version:   VER_VALUE,
		Length:    (uint32)(len(data)),
	}

	var err error = nil
	err = binary.Write(c.io, binary.BigEndian, &h.Signature)
	if nil != err {
		return err
	}

	err = binary.Write(c.io, binary.BigEndian, &h.Version)
	if nil != err {
		return err
	}

	err = binary.Write(c.io, binary.BigEndian, &h.Length)
	if nil != err {
		return err
	}

	err = binary.Write(c.io, binary.BigEndian, data)
	return err
}
