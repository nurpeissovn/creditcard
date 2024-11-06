package main

import (
	"fmt"
	"os"
	"strconv"
)

func isValid(str string) bool {
	count := 0

	for i, k := range str {
		if k == '*' && i > 11 {
			count++
		} else if k == '*' && i <= 11 {
			os.Exit(1)
		}
	}
	if count > 4 {
		os.Exit(1)
	}

	if len(str) < 13 || len(str) > 16 {
		return false
	}

	return true
}

func calculate(str string) bool {
	res := 0
	for i, k := range str {
		x, _ := strconv.Atoi(string(k))
		if i%2 == 0 {
			x *= 2
			if x > 9 {
				res += x / 10
				res += x % 10

			} else {
				res += x
			}
		} else {
			res += x
		}
	}

	if res%10 == 0 {
		return true
	} else {
		return false
	}
}

func generate(str string) {
	res := ""
	san := 0
	pow := 1
	for _, k := range str {
		if k == '*' {
			san++
			pow *= 10
		} else {
			res += string(k)
		}
	}

	for i := 0; i < pow; i++ {
		lastDigit := fmt.Sprintf("%0"+strconv.Itoa(san)+"d", i)

		if calculate(res + lastDigit) {
			fmt.Println(res + lastDigit)
		}
	}
}

func main() {
	input := os.Args[1:]

	if input[0] == "validate" {
		for _, k := range input[1:] {
			if isValid(k) && calculate(k) {
				fmt.Println("OK")
			} else {
				fmt.Println("INCORRECT")
			}
		}
	} else if input[0] == "generate" && input[1] == "--pick" {
		fmt.Println(5)
	} else if input[0] == "generate" {
		for _, k := range input[1:] {
			if isValid(k) {
				generate(string(k))
			}
		}
	}
}
