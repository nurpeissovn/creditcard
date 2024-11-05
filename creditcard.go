package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isValid(str string) bool {
	card, err := strconv.Atoi(str)
	if err != nil {
		os.Exit(1)
	}
	if len(str) < 13 {
		return false
	}
	if card == 123 {
		return false
	}
	return true
}

func main() {
	input := os.Args[1:]
	fmt.Println(os.Args[0:3])
	// if os.Args[0] == "echo" {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c := scanner.Text()
		fmt.Println(c, "check")
	}

	// }

	if input[0] == "validate" {
		for _, k := range input[1:] {
			if isValid(k) {
				fmt.Println("OK")
			} else {
				fmt.Println("INCORRECT")
			}
		}
	}
}
