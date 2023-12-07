package Day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

func (p Point) addPoint(d Point) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func getSymbols(file []byte) map[Point]string {
	symbols := map[Point]string{}

	for y, s := range strings.Fields(string(file)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				symbols[Point{x, y}] = string(r)
			}
		}
	}
	return symbols
}

func getEngineParts(file []byte, symbols map[Point]string) map[Point][]int {
	//maps points (keys) to list of ints (values)
	engineParts := map[Point][]int{}
	re := regexp.MustCompile(`\d+`)
	//all the directions in which the engine parts can be, which is one unit in any direction
	//omit 0,0 because that's the start point
	directions := []Point{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for y, s := range strings.Fields(string(file)) {
		//-1 is for all matches as opposed to a subset? - weird to not have that just be an optional param
		//find beginning and end indexes of all numbers
		for _, match := range re.FindAllStringIndex(s, -1) {
			keys := map[Point]bool{}
			for x := match[0]; x < match[1]; x++ {
				//map to points based on directions
				for _, d := range directions {
					keys[Point{x, y}.addPoint(d)] = true
				}
			}

			n, _ := strconv.Atoi(s[match[0]:match[1]])
			for p := range keys {
				if _, exists := symbols[p]; exists {
					engineParts[p] = append(engineParts[p], n)
				}
			}
		}
	}
	return engineParts

}
func part1(file []byte) int {
	symbols := getSymbols(file)
	engineParts := getEngineParts(file, symbols)

	partNumbers := 0
	for _, values := range engineParts {
		for _, value := range values {
			partNumbers += value
		}
	}
	return partNumbers
}

func part2(file []byte) int {
	symbols := getSymbols(file)
	engineParts := getEngineParts(file, symbols)
	gearRatio := 0
	for index, neighbors := range engineParts {
		if symbols[index] == "*" && len(neighbors) == 2 {
			gearRatio += neighbors[0] * neighbors[1]
		}
	}
	return gearRatio
}

func RunDay3(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	sumPart1 := part1(file)
	prodPart2 := part2(file)

	fmt.Println("Part 1: ", sumPart1)
	fmt.Println("Part 2: ", prodPart2)
}
