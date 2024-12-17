package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../day5.txt")
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

	var adjacencyList map[int][]int
	adjacencyList = make(map[int][]int)

	firstSection := true
	total := 0
	for _, line := range lines {
		if line == "" {
			firstSection = false
			continue
		}

		if firstSection {
			splitString := strings.Split(line, "|")
			key, _ := strconv.Atoi(splitString[0])
			prereq, _ := strconv.Atoi(splitString[1])

			adjacencyList[key] = append(adjacencyList[key], prereq)
			if _, ok := adjacencyList[prereq]; !ok {
				adjacencyList[prereq] = make([]int, 0)
			}
		} else {
			splitString := strings.Split(line, ",")
			intList := make([]int, len(splitString))
			for i, str := range splitString {
				num, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					continue
				}
				intList[i] = num
			}
			if !checkOrder(adjacencyList, intList) {
				sortedOrder := make([]int, 0)
				vertexSet := make(map[int]struct{})
				for _, vertex := range intList {
					vertexSet[vertex] = struct{}{}
				}
				orderUpdates(adjacencyList, intList, vertexSet, make(map[int]struct{}), &sortedOrder)
				mid := len(intList) / 2
				total += sortedOrder[mid]
			}
		}
	}

	fmt.Println(total)
}

func orderUpdates(adjacencyList map[int][]int, vertexList []int, vertexSet map[int]struct{}, done map[int]struct{}, retList *[]int) {
	for _, vertex := range vertexList {
		if _, ok := vertexSet[vertex]; !ok {
			continue
		}
		if _, ok := done[vertex]; !ok {
			done[vertex] = struct{}{}
			orderUpdates(adjacencyList, adjacencyList[vertex], vertexSet, done, retList)
			*retList = append(*retList, vertex)
		}
	}
}

func checkOrder(adjacencyList map[int][]int, partialOrder []int) bool {
	seenSet := make(map[int]struct{})
	for _, node := range partialOrder {
		for _, n := range adjacencyList[node] {
			if _, ok := seenSet[n]; ok {
				return false
			}
		}
		seenSet[node] = struct{}{}
	}
	return true
}
