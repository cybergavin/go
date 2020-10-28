package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an odd number from 1 to 25 : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	starNum, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if starNum < 1 || starNum > 25 || starNum%2 == 0 {
		err := fmt.Sprintf("%v is NOT an odd number from 1 to 25", starNum)
		log.Fatal(err)
	}
	fmt.Printf("\nHere's your %d-STAR pyramid!\n", starNum)
	for starLine := 1; starLine <= starNum; starLine += 2 {
		blankNum := (starNum - starLine) / 2
		fmt.Printf("%v%v\n", strings.Repeat(" ", blankNum), strings.Repeat("*", starLine))
	}
}
