package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	data, err := os.ReadFile("../day3.txt")
	if err != nil {
		fmt.Println("Failed to parse file")
		return
	}
	fileContent := string(data)
	contentLength := len(fileContent)
	fileContent = fileContent + "######################"

	i := 0
	total := 0
	enabled := true
	for i < contentLength {
		doCheck := fileContent[i : i+4]
		if doCheck == "do()" {
			enabled = true
			i += 4
		}
		dontCheck := fileContent[i : i+7]
		if dontCheck == "don't()" {
			enabled = false
			i += 7
		}
		if enabled {
			mulCheck := fileContent[i : i+3]
			if mulCheck == "mul" {
				i += 3
				if fileContent[i] == '(' {
					i++
					digit1 := 0
					j := 0
					flag := false
					for j < 3 && fileContent[i] != ',' {
						if unicode.IsDigit(rune(fileContent[i])) {
							digit1 = digit1*10 + int(fileContent[i]-'0')
						} else {
							flag = true
							break
						}
						i++
						j++
					}
					i++
					if flag {
						continue
					}
					digit2 := 0
					j = 0
					flag = false
					for j < 3 && fileContent[i] != ')' {
						if unicode.IsDigit(rune(fileContent[i])) {
							digit2 = digit2*10 + int(fileContent[i]-'0')
						} else {
							flag = true
							break
						}
						i++
						j++
					}
					if !flag && fileContent[i] == ')' {
						total += digit1 * digit2
					}
				}
			} else {
				i++
			}
		} else {
			i++
		}
	}
	fmt.Println(total)
}
