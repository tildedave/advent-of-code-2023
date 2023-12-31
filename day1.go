package main

import (
	"bufio"
	"fmt"
	"os"
)

func isDigit(b byte) bool {
	return 48 <= b && b <= 57
}

func getDigit(s string) int {
	// possibility one, start of string is an actual digit.
	if isDigit(s[0]) {
		return int(s[0]) - 48
	}
	// possibility two, it's a spelled out word.
	spelledOutDigits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for d := 0; d < len(spelledOutDigits); d++ {
		englishDigit := spelledOutDigits[d]
		if len(s) >= len(englishDigit) {
			if s[0:len(englishDigit)] == englishDigit {
				return d
			}
		}
	}
	return -1
}

func day1(f *os.File) {
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		var firstDigit int
		var secondDigit int

		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			firstDigit = getDigit(line[i:])
			if firstDigit != -1 {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			secondDigit = getDigit(line[i:])
			if secondDigit != -1 {
				break
			}
		}
		sum += firstDigit*10 + secondDigit
	}
	fmt.Println(sum)
}
