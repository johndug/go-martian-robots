package main

import (
	"fmt"
	"go-martian-robots-app/models"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readFile()
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}

	// parse grid size
	w, h, err := parseGridSize(lines)
	if err != nil {
		fmt.Println("Error parsing grid size")
		panic(err)
	}
	grid := models.Grid{
		Width:  w,
		Height: h,
	}

	// parse robots
	for i := 1; i < len(lines); i += 2 {
		// if line[i + 1] is empty, skip
		if i+1 >= len(lines) || lines[i+1] == "" {
			continue
		}

		x, y, facing, err := parseRobotPosition(lines[i])
		instructions := strings.Split(lines[i+1], "")

		if err != nil {
			fmt.Println("Error parsing robot position")
			panic(err)
		}

		robot := models.Robot{
			Position: models.Position{
				X: x,
				Y: y,
			},
			Facing: facing,
			Grid:   grid,
			Lost:   false,
		}

		for _, instruction := range instructions {
			if instruction != "" {
				robot.Command(instruction)
			}
		}
		fmt.Println(robot.CoordMsg())
	}
}

func readFile() ([]string, error) {
	// if os.args[1] is not set, use input.txt
	var fileName string
	if len(os.Args) < 2 {
		fileName = "input.txt"
	} else {
		fileName = os.Args[1]
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// Split by newlines and filter empty lines
	lines := strings.Split(string(content), "\n")
	var result []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return result, nil
}

func parseGridSize(lines []string) (int, int, error) {
	gridSize := strings.Split(lines[0], " ")
	w, err := strconv.Atoi(gridSize[0])
	if err != nil {
		return 0, 0, err
	}
	h, err := strconv.Atoi(gridSize[1])
	if err != nil {
		return 0, 0, err
	}
	return w, h, nil
}

func parseRobotPosition(line string) (int, int, string, error) {
	robotPosition := strings.Split(line, " ")
	x, err := strconv.Atoi(robotPosition[0])
	if err != nil {
		return 0, 0, "", err
	}
	y, err := strconv.Atoi(robotPosition[1])
	if err != nil {
		return 0, 0, "", err
	}
	return x, y, robotPosition[2], nil
}
