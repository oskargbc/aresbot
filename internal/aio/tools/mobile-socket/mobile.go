package mobile_socket

import (
	"bufio"
	"fmt"
	"net"
)

func NewMobileSocket() MobileSocket {
	return MobileSocket{}
}

type MobileSocket struct {
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func (s MobileSocket) Run() {
	conn, err := net.Dial(CONN_TYPE, "localhost:8080/ws")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn.Write([]byte("Hello from Firestorm"))

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("> " + msg)
	}
}
