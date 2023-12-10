package utils

import (
	"bufio"
	"os"
	"strconv"
)

func IsNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}



func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadLinesToInts(path string) []int {
	lines, _ := ReadLines(path)

	numbers := make([]int, len(lines))
	for i, v := range lines {
		num, _ := strconv.Atoi(v)
		numbers[i] = num
	}

	return numbers
}
