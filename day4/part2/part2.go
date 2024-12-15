package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../day4.txt")
	if err != nil {
		fmt.Println("Failed to open file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	count := 0

	for i, s := range lines[1 : len(lines)-1] {
		i++
		for j, c := range s[1 : len(s)-1] {
			j++
			if c == 'A' {
				topLeft := lines[i-1][j-1]
				topRight := lines[i+1][j-1]
				bottomLeft := lines[i-1][j+1]
				bottomRight := lines[i+1][j+1]
				if topRight == 'M' && topLeft == 'M' && bottomRight == 'S' && bottomLeft == 'S' {
					count++
				} else if topRight == 'M' && topLeft == 'S' && bottomRight == 'M' && bottomLeft == 'S' {
					count++
				} else if topRight == 'S' && topLeft == 'S' && bottomRight == 'M' && bottomLeft == 'M' {
					count++
				} else if topRight == 'S' && topLeft == 'M' && bottomRight == 'S' && bottomLeft == 'M' {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func dfs(i int, j int, matrix []string, word string, direction []int) int {
	if word == "" {
		return 1
	}
	newX := i + direction[0]
	newY := j + direction[1]
	if newX < 0 || newX >= len(matrix) || newY < 0 || newY >= len(matrix[0]) {
		return 0
	}

	if matrix[newX][newY] == word[0] {
		return dfs(newX, newY, matrix, word[1:], direction)
	}
	return 0
}
