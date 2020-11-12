package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) != 4 {
		fmt.Printf("ERROR : INVALID USAGE\n\nUSAGE: socketTest <protocol> <host:port> <timeout in seconds>\n\twhere, <protocol> = tcp|udp\n\nEXAMPLE: socketTest tcp 10.1.1.1:443 5\n")
		return
	}
	protocol := string(arguments[1])
	remoteSocket := string(arguments[2])
	t, _ := strconv.Atoi(arguments[3])
	timeout := time.Duration(t) * time.Second
	_, err := net.DialTimeout(protocol, remoteSocket, timeout)
	if err != nil {
		fmt.Printf("False")
		return
	}
	fmt.Printf("True")
}
