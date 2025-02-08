package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Robot struct {
	pos Coord
	vel Coord
}

func main() {
	file, err := os.Open("../day14.txt")
	if err != nil {
		fmt.Println("Failed to parse file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var robotList []Robot

	for scanner.Scan() {
		line := scanner.Text()
		splitString := strings.Split(line, " ")

		posString := strings.Split(splitString[0], ",")
		velString := strings.Split(splitString[1], ",")

		pos0, _ := strconv.Atoi(posString[0][2:])
		pos1, _ := strconv.Atoi(posString[1])
		vel0, _ := strconv.Atoi(velString[0][2:])
		vel1, _ := strconv.Atoi(velString[1])

		robot := Robot{
			pos: Coord{pos0, pos1},
			vel: Coord{vel0, vel1},
		}
		robotList = append(robotList, robot)
	}

	rows := 103
	cols := 101

	count := 1
	keep := true
	for {
		if !keep {
			fmt.Print(keep)
			reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')
			keep = true
		}

		for i := 0; i < len(robotList); i++ {
			robot := &robotList[i]
			robot.pos.x = mod(robot.pos.x+robot.vel.x, cols)
			robot.pos.y = mod(robot.pos.y+robot.vel.y, rows)
		}

		grid := make([][]int, rows)
		for i := range grid {
			grid[i] = make([]int, cols)
		}
		for _, robot := range robotList {
			grid[robot.pos.y][robot.pos.x]++
		}
		if checkClusters(grid, rows, cols) {
			printRobots(grid)
			fmt.Println()
			fmt.Print(count)
			fmt.Println()
			keep = false
		}
		count++

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading the file")
		}
	}
}

func printRobots(grid [][]int) {
	for i := 0; i < len(grid); i++ { // Iterate over rows
		for j := 0; j < len(grid[i]); j++ { // Iterate over columns
			if grid[i][j] == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Println() // Newline after each row
	}
}

func checkClusters(grid [][]int, maxRows int, maxCols int) bool {
	visited := make(map[Coord]struct{})
	maxCount := 0
	for i := 0; i < len(grid); i++ { // Iterate over rows
		for j := 0; j < len(grid[i]); j++ { // Iterate over columns
			if grid[i][j] != 0 {
				count := dfs(grid, Coord{i, j}, maxRows, maxCols, visited)
				if count > maxCount {
					maxCount = count
				}
			}
		}
	}

	if maxCount > 100 {
		return true
	} else {
		return false
	}
}

func dfs(grid [][]int, coord Coord, maxRows int, maxCols int, visited map[Coord]struct{}) int {
	if _, ok := visited[coord]; ok {
		return 0
	}
	if coord.x < 0 || coord.x >= maxRows || coord.y < 0 || coord.y >= maxCols || grid[coord.x][coord.y] == 0 {
		return 0
	}
	visited[coord] = struct{}{}
	directions := [8][2]int{
		{0, 1},   // North
		{1, 1},   // NorthEast
		{1, 0},   // East
		{1, -1},  // SouthEast
		{0, -1},  // South
		{-1, -1}, // SouthWest
		{-1, 0},  // West
		{-1, 1},  // NorthWest
	}

	count := 0
	for i := 0; i < len(directions); i++ {
		newX := coord.x + directions[i][0]
		newY := coord.y + directions[i][1]
		count += dfs(grid, Coord{newX, newY}, maxRows, maxCols, visited)
	}

	return count + 1
}

func mod(a, b int) int {
	return (a%b + b) % b
}
