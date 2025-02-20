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

	var quadrants [4]int
	for _, robot := range robotList {
		newX := mod(robot.pos.x+(robot.vel.x*100), cols)
		newY := mod(robot.pos.y+(robot.vel.y*100), rows)

		rowMiddle := rows / 2
		colMiddle := cols / 2

		if newX > colMiddle && newY > rowMiddle {
			quadrants[0]++
		} else if newX > colMiddle && newY < rowMiddle {
			quadrants[1]++
		} else if newX < colMiddle && newY > rowMiddle {
			quadrants[2]++
		} else if newX < colMiddle && newY < rowMiddle {
			quadrants[3]++
		}
	}

	result := 1
	for _, num := range quadrants {
		result *= num
	}
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file")
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}
