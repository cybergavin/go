/*
Author: cybergavin
Creation Date: 11th November 2020
Description: The socketTest program tests connectivity to a socket and returns "True" or "False" based on whether there is connectivity.
*/
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
	helpMessage := `
	USAGE: socketTest <protocol> <host:port> <timeout in seconds>
		   where, <protocol> = tcp|udp
	EXAMPLE: socketTest tcp 10.1.1.1:443 5
	REFERENCE: https://github.com/cybergavin/go/blob/master/sandbox/socketTest/main.go				
	`
	if len(arguments) != 4 {
		fmt.Println(helpMessage)
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
