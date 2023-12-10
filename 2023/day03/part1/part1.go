package main

import (
	"fmt"
	"regexp"
	"solution/utils"
	"strconv"
	"unicode"
)

func isSymbol(c rune) bool {
    return c != '.' && !unicode.IsNumber(c) && !unicode.IsLetter(c)
}


type Point struct {
    X int 
    Y int 
}




func populateSymbolLocations(lines []string) map[Point]bool {
    pointSet := map[Point]bool{}

    for x, line := range lines {
        for y, c := range line {
            if isSymbol(c) {
                pointSet[Point{x,y}] = true
            }
        }
    }
    return pointSet
}


func findPartNumbers(lines []string) map[Point]string {
    partNumberLocations := map[Point]string{}
	re := regexp.MustCompile(`\d+`)

    for x, line := range lines {
        matches := re.FindAllStringIndex(line, -1)
        
        for _, match := range matches {
            startIndex := match[0]
            partNumberLocations[Point{x, startIndex}] = line[match[0]:match[1]]
        }
    }

	return partNumberLocations
}


func checkForAdjacentSymbol(p Point, symbols map[Point]bool) bool {
    if symbols[Point{p.X-1, p.Y-1}] {
        return true
    }
    if symbols[Point{p.X-1, p.Y}] {
        return true
    }
    if symbols[Point{p.X-1, p.Y+1}] {
        return true
    }
    if symbols[Point{p.X, p.Y-1}] {
        return true
    }
    if symbols[Point{p.X, p.Y+1}] {
        return true
    }
    if symbols[Point{p.X+1, p.Y-1}] {
        return true
    }
    if symbols[Point{p.X+1, p.Y}] {
        return true
    }
    if symbols[Point{p.X+1, p.Y+1}] {
        return true
    }
    return false
}



func compute(input []string) int {
    ret := 0

    symbolLocations := populateSymbolLocations(input)
    partNumberLocations := findPartNumbers(input)

    for point, partNo := range partNumberLocations {
        isAddNumber := false

        for idx, _ := range partNo {
            if isAddNumber == true {
                break 
            }

            numberPoint := Point{point.X, point.Y+idx}

            if checkForAdjacentSymbol(numberPoint, symbolLocations) == true {
                isAddNumber = true
                num, _ := strconv.Atoi(partNo)
                ret += num
            }
        }
    }

    return ret
}

func main() {
    lines , _ := utils.ReadLines("../input.txt")
    fmt.Println(compute(lines))
}
