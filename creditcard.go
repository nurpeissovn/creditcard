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
		if k == '*' {
			for _, l := range str[i:] {
				if l != '*' {
					return false
				}
			}
		}
		if str[:2] == "34" && len(str) != 15 || str[:2] == "37" && len(str) != 15 {
			return false
		}

	}
	return true
}

func calculate(str string) bool {
	res := 0
	for i, k := range str {
		x, err := strconv.Atoi(string(k))
		if err != nil {
			return false
		}
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
	if san == 0 {
		os.Exit(1)
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

func readFile(fileName string) map[string]string {
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
	Map := readFile("issuers.txt")
	Mapb := readFile("brands.txt")
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

func stdInput(check bool) []string {
	scanner := bufio.NewScanner(os.Stdin)
	file, _ := os.Stdin.Stat()
	var input []string = []string{"", ""}
	if (file.Mode() & os.ModeCharDevice) != 0 {
		os.Exit(1)
	}
	if check {
		input = append(input, "")
	}
	for scanner.Scan() {

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
	return input
}

func main() {
	input := os.Args[1:]

	if len(input) < 2 || input[0] == "information" && len(input) < 4 {
		os.Exit(1)
	}

	if input[0] == "validate" && isValid(input[1]) || input[0] == "validate" && input[1] == "--stdin" {

		if input[1] == "--stdin" {
			if len(input) > 2 {
				os.Exit(1)
			}
			input = stdInput(false)
		}

		for _, k := range input[1:] {
			if isValid(k) && calculate(k) && k != "" {
				fmt.Println("OK")
			} else if k != "" {
				fmt.Println("INCORRECT")
				os.Exit(1)
			}
		}
	} else if len(input) < 4 && len(input) > 2 && input[0] == "generate" && input[1] == "--pick" && isValid(input[2]) {
		generate(input[2], false)
	} else if input[0] == "generate" && len(input) < 3 {
		for _, k := range input[1:] {
			if isValid(k) {
				generate(string(k), true)
			} else {
				os.Exit(1)
			}
		}
	} else if input[0] == "information" && input[1] == "--brands=brands.txt" && input[2] == "--issuers=issuers.txt" && len(input) > 3 {

		if input[3] == "--stdin" {
			if len(input) > 4 {
				os.Exit(1)
			}
			input = stdInput(true)
		} else if file, _ := os.Stdin.Stat(); (file.Mode()&os.ModeCharDevice) == 0 && input[3] != "--stdin" {
			os.Exit(1)
		}
		fin := 2
		for _, k := range input[3:] {
			if k[0] == '4' {
				fin = 1
			}
			if isValid(k) && calculate(k) && k != "" {
				bName := readFile("brands.txt")
				issuerNmae := readFile("issuers.txt")
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

		if len(brand) >= 9 && len(issuer) >= 10 {
			issueCard(brand[8:], issuer[9:])
		} else {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}
