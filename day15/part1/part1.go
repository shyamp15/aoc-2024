package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

func addPos(a Pos, b Pos) Pos {
	return Pos{a.x + b.x, a.y + b.y}
}

type Obj int

const (
	Wall Obj = iota
	Empty
	Box
	Robot
)

type Movement int

const (
	Left Movement = iota
	Right
	Up
	Down
)

func main() {
	file, err := os.Open("../day15.txt")
	if err != nil {
		fmt.Println("Failed to parse file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstPart bool = true
	var grid [][]Obj
	var robotPos Pos
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			firstPart = false
		} else if firstPart {
			parseRow(line, &grid, &robotPos)
		} else {
			doMovements(line, grid, &robotPos)
		}
	}
	fmt.Println(calcScore(grid))
	printBoard(grid)
}

func parseRow(line string, grid *[][]Obj, robotPos *Pos) {
	var row []Obj
	for i, c := range line {
		if c == '#' {
			row = append(row, Wall)
		} else if c == '.' {
			row = append(row, Empty)
		} else if c == 'O' {
			row = append(row, Box)
		} else if c == '@' {
			row = append(row, Robot)
			robotPos.x = len(*grid)
			robotPos.y = i
		}
	}
	*grid = append(*grid, row)
}

func doMovements(line string, grid [][]Obj, robotPos *Pos) {
	for _, c := range line {
		if c == '<' {
			doMovement(Pos{0, -1}, robotPos, grid)
		} else if c == '>' {
			doMovement(Pos{0, 1}, robotPos, grid)
		} else if c == '^' {
			doMovement(Pos{-1, 0}, robotPos, grid)
		} else if c == 'v' {
			doMovement(Pos{1, 0}, robotPos, grid)
		}
	}
}

func doMovement(move Pos, robotPos *Pos, grid [][]Obj) {
	newRobotPos := addPos(move, *robotPos)
	switch grid[newRobotPos.x][newRobotPos.y] {
	case Wall:
		//Do nothing
		return
	case Empty:
		grid[robotPos.x][robotPos.y] = Empty
		grid[newRobotPos.x][newRobotPos.y] = Robot
		robotPos.x = robotPos.x + move.x
		robotPos.y = robotPos.y + move.y
	case Box:
		newPos := addPos(newRobotPos, move)
		numBoxes := 1
		for grid[newPos.x][newPos.y] == Box {
			numBoxes++
			newPos = addPos(newPos, move)
		}
		if grid[newPos.x][newPos.y] == Wall {
			//We have a line of boxes that hit a wall so nothing can move so just return
			return
		}
		currPos := newRobotPos
		grid[robotPos.x][robotPos.y] = Empty
		grid[currPos.x][currPos.y] = Robot
		for i := 0; i < numBoxes; i++ {
			currPos = addPos(currPos, move)
			grid[currPos.x][currPos.y] = Box
		}
		robotPos.x = robotPos.x + move.x
		robotPos.y = robotPos.y + move.y
	default:
		fmt.Println("Error - shouldn't be able to move into a robot slot")
	}
}

func printBoard(grid [][]Obj) {
	for i, _ := range grid {
		for j, _ := range grid[i] {
			switch grid[i][j] {
			case Wall:
				fmt.Print("#")
			case Empty:
				fmt.Print(".")
			case Box:
				fmt.Print("O")
			case Robot:
				fmt.Print("@")
			}
		}
		fmt.Println("")
	}
}

func calcScore(grid [][]Obj) int {
	score := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == Box {
				score += i*100 + j
			}
		}
	}
	return score
}
