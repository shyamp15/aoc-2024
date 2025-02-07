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

	for _, robot := range robotList {
		newX := robot.pos.x + ((robot.vel.x * 100) % rows)
		newY := robot.pos.y + ((robot.vel.y * 100) % cols)
		fmt.Println(newX, newY)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file")
	}
}
