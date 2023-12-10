package main

import (
	"fmt"
	"solution/utils"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var numberMap = map[string]string{
	"one":   "1",
    "two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Entry struct {
	position int
	value    string
}

func NewEntry(position int, value string) Entry {
	return Entry{
		position,
		value,
	}
}

func firstAndLastNumber(line string) string {
	numberStrings := []Entry{}

    // Populate with spelled out numbers 
	for key := range numberMap {
        value := numberMap[key]
		idx_first := strings.Index(line, key)
		idx_last := strings.LastIndex(line, key)

		if idx_first != -1 {
			numberStrings =
				append(numberStrings, NewEntry(idx_first, value))
		}
        if idx_last != -1 && idx_last != idx_first {
			numberStrings =
				append(numberStrings, NewEntry(idx_last, value))
        }
	}

    // populate list with char numbers
    for idx, c := range line {
        if unicode.IsDigit(c) {
            numberStrings = append(numberStrings, NewEntry(idx, string(c)))
        }
    }

	sort.Slice(numberStrings, func(i, j int) bool {
		return numberStrings[i].position < numberStrings[j].position
	})

	element1 := numberStrings[0]

	if len(numberStrings) == 1 {
		return element1.value + element1.value
	}

	element2 := numberStrings[len(numberStrings)-1]
	return element1.value + element2.value
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
	lines, _ := utils.ReadLines("./input.txt")
	fmt.Println(compute(lines))
}
