package main

import (
	"fmt"
	"math"
	"regexp"
	"solution/utils"
	"strconv"
)


type Point struct {
    X int 
    Y int 
}

type Num struct {
    Val        int 
    StartPoint Point
    len        int 
}

func NewNum(number string, x int, y int) Num {
    num, _ := strconv.Atoi(number)

    return Num{Val: num, StartPoint: Point{x,y}, len: len(number)}
}


func arePointsAdjacent(p1, p2 Point) bool {
	deltaX := math.Abs(float64(p1.X - p2.X))
	deltaY := math.Abs(float64(p1.Y - p2.Y))

	// common edge
	if deltaX == 0 && deltaY == 1 || deltaX == 1 && deltaY == 0 {
		return true
	}

	// common corner
	if deltaX == 1 && deltaY == 1 {
		return true
	}

	return false
}


func (n Num)Adjacent (p Point) bool {
    for y := 0; y < n.len; y++ {

        if arePointsAdjacent(Point{n.StartPoint.X, y + n.StartPoint.Y}, p) {
            return true
        }
    }

    return false
}


func findPartNumbers(lines []string) []Num{
    numbers := []Num{}
 	re := regexp.MustCompile(`\d+`)

    for x, line := range lines {
        matches := re.FindAllStringIndex(line, -1)
        
        for _, match := range matches {
            startIndex := match[0]
            number := line[match[0]:match[1]]
            numbers = append(numbers, NewNum(number, x, startIndex))
        }
    }

	return numbers
}



func createGearArray(lines []string) []Point {
    gears := []Point{}

    for x, line := range lines {
        for y, c := range line {
            if c == '*' {
                gears = append(gears, Point{x,y})
            }
        }
    }
    return gears
}


func compute(input []string) int {
    ret := 0

    gears := createGearArray(input)
    numbers := findPartNumbers(input)
    
    for _, point := range gears {
        adjacent := []Num{}
        
        for _, num := range numbers {
            if num.Adjacent(point) {
                adjacent = append(adjacent, num)
            }
        }

        if len(adjacent) == 2 {
            ret += (adjacent[0].Val * adjacent[1].Val)
        }
    }


    return ret
}

func main() {
    lines , _ := utils.ReadLines("../input.txt")
    fmt.Println(compute(lines))
}

