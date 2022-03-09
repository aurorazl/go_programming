package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_programming/chatRoom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	read, err := t.Conn.Read(t.Buf[:4])
	if read != 4 || err != nil {
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("recv pkg error", err)
		return
	}
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	write, err := t.Conn.Write(t.Buf[:4])
	if write != 4 || err != nil {
		fmt.Println("send error", err)
		return
	}
	_, err = t.Conn.Write(data)
	if err != nil {
		fmt.Println("send data error ", err)
		return
	}
	return
}
