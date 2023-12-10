package main

import (
	"fmt"
	"os"
	"regexp"
	"solution/utils"
	"strconv"
	"strings"
)

var game_re = regexp.MustCompile(`Game (\d+)`)
var red_re = regexp.MustCompile(`(\d+)\s*red`)
var green_re = regexp.MustCompile(`(\d+)\s*green`)
var blue_re = regexp.MustCompile(`(\d+)\s*blue`)

type GameSet struct {
	Red  int
	Green int
	Blue int
}

var MaxPieces = GameSet{
    Red: 12,
    Green: 13,
    Blue: 14,
}

func (gs *GameSet) IsLessThan(other GameSet) bool {
    if gs.Red > other.Red || gs.Green > other.Green || gs.Blue > other.Blue {
        return false
    }
    return true
}


func isGamePossible(line string) int {

    gameNumberMatches := game_re.FindStringSubmatch(line)
    if len(gameNumberMatches) < 1 {
        os.Exit(1)
    }

    gameNum, _ := strconv.Atoi(gameNumberMatches[1])
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

        if !gameSet.IsLessThan(MaxPieces) {
            return 0
        }
    }

    return gameNum
}



func compute(input []string) int {
    ret := 0
    for _, line := range input {
        ret += isGamePossible(line)
    }
    return ret
}

func main() {
    lines , _ := utils.ReadLines("../input.txt")
    fmt.Println(compute(lines))
}
