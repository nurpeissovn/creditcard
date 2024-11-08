package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func isValid(str string) bool {
	for i, k := range str {
		if k == '*' && len(str)-i > 4 || k < 48 && k != '*' || k > 57 && k != '*' || len(str) < 13 || len(str) > 16 {
			return false
		}
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

	return res%10 == 0
}

func generate(str string, check bool) {
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
		PickFlag := (res + fmt.Sprintf("%0"+strconv.Itoa(san)+"d", rand.Intn(pow)))

		if !check && calculate(PickFlag) {
			fmt.Println(PickFlag)
			return
		} else if calculate(res+lastDigit) && check {
			fmt.Println(res + lastDigit)
		}
	}
}

func readFfile(fileName string) map[string]string {
	file, err := os.Open(fileName)
	Map := make(map[string]string)
	brandName := ""
	preNum := ""
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, k := range scanner.Text() {
			if k >= 48 && k <= 57 {
				preNum += string(k)
			} else if k != ':' {
				brandName += string(k)
			}
		}

		Map[preNum] = brandName
		preNum = ""
		brandName = ""
	}

	return Map
}

func issueCard(brand, issuer string) {
	value := ""
	check := ""
	Map := readFfile("issuers.txt")
	Mapb := readFfile("brands.txt")
	for key := range Mapb {
		if Mapb[key] == brand {
			check = string(key[0])
		}
	}
	for key := range Map {
		if Map[key] == issuer && check == string(key[0]) {
			value = key
		}
	}

	if len(value) == 0 {
		os.Exit(1)
	}

	for {
		card := value + fmt.Sprintf("%010d", rand.Intn(1000000000))
		if calculate(card) {
			fmt.Println(card)
			return
		}
	}
}

func main() {
	input := os.Args[1:]

	if len(input) < 2 {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)

	if input[1] == "--stdin" { // validate --stdin not working

		if len(input) != 2 {
			os.Exit(1)
		}
		for scanner.Scan() {
			input[1] = ""
			var temp string

			for i, k := range scanner.Text() {
				if k != ' ' {
					temp += string(k)
				}
				if k == ' ' || len(scanner.Text())-1 == i {
					input = append(input, temp)

					temp = ""
				}
			}
		}

	}
	if input[0] == "validate" && isValid(input[1]) || input[0] == "validate" && input[1] == "--stdin" {
		for _, k := range input[1:] {
			if isValid(k) && calculate(k) && k != "" {
				fmt.Println("OK")
			} else if k != "" {
				fmt.Println("INCORRECT")
			}
		}
	} else if len(input) < 4 && input[0] == "generate" && input[1] == "--pick" && isValid(input[2]) {
		generate(input[2], false)
	} else if input[0] == "generate" && len(input) < 3 {
		for _, k := range input[1:] {
			if isValid(k) {
				generate(string(k), true)
			} else {
				os.Exit(1)
			}
		}
	} else if input[0] == "information" && input[1] == "--brands=brands.txt" && input[2] == "--issuers=issuers.txt" {
		fin := 2
		for _, k := range input[3:] {
			if k[0] == '4' {
				fin = 1
			}
			if isValid(k) && calculate(k) {
				bName := readFfile("brands.txt")
				issuerNmae := readFfile("issuers.txt")
				fmt.Println(k)
				fmt.Println("Correct: " + "yes")
				fmt.Println("Card Brand: " + bName[k[:fin]])
				fmt.Println("Card Issuer: " + issuerNmae[k[:6]])

			} else {
				fmt.Println(k)
				fmt.Println("Correct: " + "No")
				fmt.Println("Card Brand: " + "-")
				fmt.Println("Card Issuer: " + "-")

			}
		}
	} else if input[0] == "issue" && input[1] == "--brands=brands.txt" && input[2] == "--issuers=issuers.txt" && len(input) == 5 {
		brand := input[3]
		issuer := input[4]

		if len(brand) >= 12 && len(issuer) >= 18 {
			issueCard(brand[8:], issuer[9:])
		} else {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}
