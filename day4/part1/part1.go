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
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for i, s := range lines {
		for j, c := range s {
			if c == 'X' {
				for _, direction := range directions {
					count += dfs(i, j, lines, "MAS", direction)
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
