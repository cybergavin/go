/*
For a user-provided height (number of rows), display a star pyramid/triangle
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Accept user input and validate
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("This program will display a star pyramid.")
	fmt.Printf("\nHow many rows do you want in the pyramid? : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rowNum, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if rowNum < 2 {
		err := "You must have atleast 2 rows for a pyramid!"
		log.Fatal(err)
	}
	// Clear screen and display output
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Printf("%v\nHere's your %d-ROW STAR pyramid!\n%v\n", strings.Repeat("-", 50), rowNum, strings.Repeat("-", 50))
	lastRowStars := 2*rowNum - 1
	for starRow := 1; starRow <= rowNum; starRow++ {
		rowStars := 2*starRow - 1
		blankNum := (lastRowStars - rowStars) / 2
		fmt.Printf("%v%v\n", strings.Repeat(" ", blankNum), strings.Repeat("*", rowStars))
	}
}
