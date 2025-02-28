package smb

import (
	"bytes"
	"context"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/projectdiscovery/nuclei/v3/pkg/js/libs/structs"
	"github.com/projectdiscovery/nuclei/v3/pkg/protocols/common/protocolstate"
)

const (
	pkt = "\x00\x00\x00\xc0\xfeSMB@\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00$\x00\x08\x00\x01\x00\x00\x00\x7f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00x\x00\x00\x00\x02\x00\x00\x00\x02\x02\x10\x02\"\x02$\x02\x00\x03\x02\x03\x10\x03\x11\x03\x00\x00\x00\x00\x01\x00&\x00\x00\x00\x00\x00\x01\x00 \x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\x00\n\x00\x00\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00"
)

// DetectSMBGhost tries to detect SMBGhost vulnerability
// by using SMBv3 compression feature.
func (c *SMBClient) DetectSMBGhost(host string, port int) (bool, error) {
	addr := net.JoinHostPort(host, strconv.Itoa(port))
	conn, err := protocolstate.Dialer.Dial(context.TODO(), "tcp", addr)
	if err != nil {
		return false, err

	}
	defer conn.Close()

	_, err = conn.Write([]byte(pkt))
	if err != nil {
		return false, err
	}

	buff := make([]byte, 4)
	nb, _ := conn.Read(buff)
	args, err := structs.Unpack(">I", buff[:nb])
	if err != nil {
		return false, err
	}
	if len(args) != 1 {
		return false, errors.New("invalid response")
	}

	length := args[0].(int)
	data := make([]byte, length)
	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err := conn.Read(data)
	if err != nil {
		return false, err
	}
	data = data[:n]

	if !bytes.Equal(data[68:70], []byte("\x11\x03")) || !bytes.Equal(data[70:72], []byte("\x02\x00")) {
		return false, nil
	}
	return true, nil
}
