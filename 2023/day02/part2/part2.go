package main

import (
	"fmt"
	"regexp"
	"solution/utils"
	"strconv"
	"strings"
)

var red_re = regexp.MustCompile(`(\d+)\s*red`)
var green_re = regexp.MustCompile(`(\d+)\s*green`)
var blue_re = regexp.MustCompile(`(\d+)\s*blue`)

type GameSet struct {
	Red  int
	Green int
	Blue int
}

func (gs *GameSet) UpdateMax(other GameSet) {
    if other.Red > gs.Red {
        gs.Red = other.Red
    }
    if other.Green > gs.Green {
        gs.Green = other.Green
    }
    if other.Blue > gs.Blue {
        gs.Blue = other.Blue
    }
}

func (gs GameSet) Power() int {
    return gs.Red * gs.Green * gs.Blue
}


func computePower(line string) int {
    globalGame := GameSet{}

    gameParts := strings.Split(line, ";")
    for _, part := range gameParts {
        gameSet := GameSet{}

        redCount := red_re.FindStringSubmatch(part)
        greenCount := green_re.FindStringSubmatch(part)
        blueCount := blue_re.FindStringSubmatch(part)

        if redCount != nil {
            reds, _ := strconv.Atoi(redCount[1])
            gameSet.Red = reds
        }
        if blueCount != nil {
            blues, _ := strconv.Atoi(blueCount[1])
            gameSet.Blue = blues
        }
        if greenCount != nil {
            greens, _ := strconv.Atoi(greenCount[1])
            gameSet.Green = greens
        }

        globalGame.UpdateMax(gameSet)
    }

    return globalGame.Power()
}



func compute(input []string) int {
    ret := 0
    for _, line := range input {
        ret += computePower(line)
    }
    return ret
}

func main() {
    lines , _ := utils.ReadLines("../input.txt")
    fmt.Println(compute(lines))
}
