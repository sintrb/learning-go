package main

import (
	"fmt"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Faltal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "172.16.0.200:9999")
	checkErr(err)
	var msg [128]byte
	msg[0] = 'a'
	msg[1] = 'b'
	msg[2] = 'c'
	msg[3] = '\r'
	msg[4] = '\n'
	msg[5] = '\r'
	msg[6] = '\n'
	msg[7] = '\r'
	msg[8] = '\n'
	_, err = conn.Write(msg[0:])

	count, err := conn.Read(msg[0:])

	fmt.Printf("len=%d %v", count, msg[0:count])
}
