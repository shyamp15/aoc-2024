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
	OpenBox
	CloseBox
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
			row = append(row, Wall)
		} else if c == '.' {
			row = append(row, Empty)
			row = append(row, Empty)
		} else if c == 'O' {
			row = append(row, OpenBox)
			row = append(row, CloseBox)
		} else if c == '@' {
			row = append(row, Robot)
			row = append(row, Empty)
			robotPos.x = len(*grid)
			robotPos.y = i * 2
		}
	}
	*grid = append(*grid, row)
}

func doMovements(line string, grid [][]Obj, robotPos *Pos) {
	for _, c := range line {
		if c == '<' {
			doMovement(Pos{0, -1}, robotPos, &grid)
		} else if c == '>' {
			doMovement(Pos{0, 1}, robotPos, &grid)
		} else if c == '^' {
			doMovement(Pos{-1, 0}, robotPos, &grid)
		} else if c == 'v' {
			doMovement(Pos{1, 0}, robotPos, &grid)
		}
	}
}

func doMovement(move Pos, robotPos *Pos, gridPtr *[][]Obj) {
	newRobotPos := addPos(move, *robotPos)
	grid := (*gridPtr)
	switch grid[newRobotPos.x][newRobotPos.y] {
	case Wall:
		//Do nothing
		return
	case Empty:
		grid[robotPos.x][robotPos.y] = Empty
		grid[newRobotPos.x][newRobotPos.y] = Robot
		robotPos.x = robotPos.x + move.x
		robotPos.y = robotPos.y + move.y
	case OpenBox:
		var openBoxes = []Pos{newRobotPos}
		var closedBoxes = []Pos{addPos(newRobotPos, Pos{0, 1})}
		var newBox = []Pos{newRobotPos, addPos(newRobotPos, Pos{0, 1})}
		openBoxPtr, closedBoxPtr, canMove := getMovableBoxes(newBox, &openBoxes, &closedBoxes, &grid, move)
		if canMove {
			moveObjects(robotPos, openBoxPtr, closedBoxPtr, gridPtr, move)
		}
	case CloseBox:
		var openBoxes = []Pos{addPos(newRobotPos, Pos{0, -1})}
		var closedBoxes = []Pos{newRobotPos}
		var newBox = []Pos{newRobotPos, addPos(newRobotPos, Pos{0, -1})}
		openBoxPtr, closedBoxPtr, canMove := getMovableBoxes(newBox, &openBoxes, &closedBoxes, &grid, move)
		if canMove {
			moveObjects(robotPos, openBoxPtr, closedBoxPtr, gridPtr, move)
		}
	default:
		fmt.Println("Error - shouldn't be able to move into a robot slot")
	}
}

func getMovableBoxes(currBox []Pos, openBoxes *[]Pos, closedBoxes *[]Pos, grid *[][]Obj, move Pos) (*[]Pos, *[]Pos, bool) {
	for _, pos := range currBox {
		newPos := addPos(pos, move)
		//If we hit a wall then we can just early return false since nothing can move
		if (*grid)[newPos.x][newPos.y] == Wall {
			return openBoxes, closedBoxes, false
		}

		canMove := true
		if (*grid)[newPos.x][newPos.y] == OpenBox {
			*openBoxes = append(*openBoxes, newPos)
			*closedBoxes = append(*closedBoxes, addPos(newPos, Pos{0, 1}))
			var newBox []Pos
			if move.x == 0 && move.y == -1 {
				newBox = []Pos{newPos}
			} else if move.x == 0 && move.y == 1 {
				newBox = []Pos{addPos(newPos, Pos{0, 1})}
			} else {
				newBox = []Pos{newPos, addPos(newPos, Pos{0, 1})}
			}
			openBoxes, closedBoxes, canMove = getMovableBoxes(newBox, openBoxes, closedBoxes, grid, move)
		} else if (*grid)[newPos.x][newPos.y] == CloseBox {
			*closedBoxes = append(*closedBoxes, newPos)
			*openBoxes = append(*openBoxes, addPos(newPos, Pos{0, -1}))
			var newBox []Pos
			if move.x == 0 && move.y == -1 {
				newBox = []Pos{addPos(newPos, Pos{0, -1})}
			} else if move.x == 0 && move.y == 1 {
				newBox = []Pos{newPos}
			} else {
				newBox = []Pos{newPos, addPos(newPos, Pos{0, -1})}
			}
			openBoxes, closedBoxes, canMove = getMovableBoxes(newBox, openBoxes, closedBoxes, grid, move)
		}
		if !canMove {
			return openBoxes, closedBoxes, false
		}
	}
	return openBoxes, closedBoxes, true
}

func moveObjects(robotPos *Pos, openBoxes *[]Pos, closedBoxes *[]Pos, gridPtr *[][]Obj, move Pos) {
	grid := *gridPtr
	hashSet := make(map[Pos]bool)

	//Move openBoxes
	for _, pos := range *openBoxes {
		currPos := addPos(pos, move)
		grid[currPos.x][currPos.y] = OpenBox
		hashSet[currPos] = true
	}
	//Move closedBoxes
	for _, pos := range *closedBoxes {
		currPos := addPos(pos, move)
		grid[currPos.x][currPos.y] = CloseBox
		hashSet[currPos] = true
	}

	//Clear old boxes
	for _, pos := range *openBoxes {
		if _, exists := hashSet[pos]; !exists {
			grid[pos.x][pos.y] = Empty
		}
	}
	for _, pos := range *closedBoxes {
		if _, exists := hashSet[pos]; !exists {
			grid[pos.x][pos.y] = Empty
		}
	}
	//Move robot
	newPos := addPos(*robotPos, move)
	grid[robotPos.x][robotPos.y] = Empty
	grid[newPos.x][newPos.y] = Robot
	robotPos.x = robotPos.x + move.x
	robotPos.y = robotPos.y + move.y
}

func arrayToHashSet(arr []Pos) map[Pos]bool {
	hashSet := make(map[Pos]bool)
	for _, val := range arr {
		hashSet[val] = true // Add element to the set
	}
	return hashSet
}

func printBoard(grid [][]Obj) {
	for i, _ := range grid {
		for j, _ := range grid[i] {
			switch grid[i][j] {
			case Wall:
				fmt.Print("#")
			case Empty:
				fmt.Print(".")
			case OpenBox:
				fmt.Print("[")
			case CloseBox:
				fmt.Print("]")
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
			if grid[i][j] == OpenBox {
				score += i*100 + j
			}
		}
	}
	return score
}
