package main

import (
	"fmt"
	"os"
	"solution/utils"
    "strconv"
	"unicode"
)


func firstAndLastNumber(line string) string {
    numericChars := []rune{}

    for _, c := range line {
        if unicode.IsDigit(c) {
            numericChars = append(numericChars, c)
        }
    }


    if len(numericChars) == 0 {
        os.Exit(1)
    }

    if len(numericChars) == 1 {
        return string(numericChars[0]) + string(numericChars[0])
    }

    return string(numericChars[0]) + string(numericChars[len(numericChars)-1])
}



func compute(input []string) int {
    ret := 0
    for _, line := range input {
        num, _ := strconv.Atoi(firstAndLastNumber(line))
        ret += num
    }
    return ret
}

func main() {
    lines , _ := utils.ReadLines("./input.txt")
    fmt.Println(compute(lines))
}
