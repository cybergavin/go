/*
Author: cybergavin
Creation Date: 11th November 2020
Description: The tcpPQ program tests TCP connectivity to a socket and returns "True" or "False" based on whether there is connectivity.
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
	USAGE: tcpPQ <host:port> [timeout in seconds]
	timeout is optional and the default timeout is 5 seconds
	EXAMPLE: tcpPQ 10.1.1.1:443 
	REFERENCE: https://github.com/cybergavin/go/blob/master/sandbox/tcpPQ/main.go				
	`
	if len(arguments) != 2 && len(arguments) != 3 {
		fmt.Println(helpMessage)
		return
	}
	var remoteSocket string
	var t int
	var err error
	for index, argument := range os.Args {
		if index == 1 {
			remoteSocket = argument
			t = 5
		}
		if index == 2 {
			t, err = strconv.Atoi(argument)
			if err != nil {
				fmt.Println(helpMessage)
				return
			}
		}
	}
	timeout := time.Duration(t) * time.Second
	conn, err := net.DialTimeout("tcp", remoteSocket, timeout)
	if err != nil {
		fmt.Printf("False")
		return
	}
	fmt.Printf("True")
	conn.Close()
}
