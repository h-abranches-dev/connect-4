package utils

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func ListSrvAddr(port int) string {
	return NewAddress("0.0.0.0", port)
}

func NewAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func CloseConn(conn *grpc.ClientConn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("error closing the connection (%s)", conn.Target())
	}
}
