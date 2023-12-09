package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Assignment struct {
    low int 
    high int 
}

func NewAssignment(r string) Assignment {
    section_range := strings.Split(r, "-")
    low, err := strconv.Atoi(section_range[0])
    if err != nil {
        panic(err)
    }
    high, err := strconv.Atoi(section_range[1])
    if err != nil {
        panic(err)
    }
    return Assignment{low: low, high: high}
}


func RangeEncapsulatesPartial(a1 *Assignment, a2 *Assignment) bool {
    if a2.low >= a1.low && a2.low <= a1.high || 
        a1.low >= a2.low && a1.low <= a2.high {
            return true
        }
    return false
}


func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    ret := 0

    for scanner.Scan() {
        ranges := strings.Split(scanner.Text(), ",")
        if len(ranges) != 2 {
            panic("invalid input")
        }

        A1 := NewAssignment(ranges[0])
        A2 := NewAssignment(ranges[1])

        if RangeEncapsulates(&A1, &A2) {
            ret += 1
        }
    }
    return ret
}

func main() {
    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}
