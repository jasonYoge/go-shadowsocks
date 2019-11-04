package go_shadowsocks

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

const BufSize = 1024

type SecureSocket struct {
	Cipher *Cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func (secureSocket *SecureSocket) DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
	n, err = conn.Read(bs)
	if err != nil {
		fmt.Println("Read from tcp connection error: ", err.Error())
		return n, err
	}

	secureSocket.Cipher.decode(bs[:n])
	return n, err
}

func (secureSocket *SecureSocket) EncodeWrite(conn *net.TCPConn, bs []byte) (int, error) {
	secureSocket.Cipher.encode(bs)
	return conn.Write(bs)
}

func (secureSocket *SecureSocket) EncodePipe(dist *net.TCPConn, src *net.TCPConn) error {
}
