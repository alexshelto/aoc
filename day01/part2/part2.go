package main 

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "sort"
)

func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))

    local_sum := 0
    calorie_sums := []int{}

    for scanner.Scan() {
        line := scanner.Text()
        // If we hit newline new group 
        if len(line) == 0 {
            calorie_sums = append(calorie_sums, local_sum)
            local_sum = 0
        } else {
            calories, err := strconv.Atoi(line)
            if err != nil {
                panic("attempted to parse invalid int")
            }
            local_sum += calories 
        }
    }

    sort.Slice(calorie_sums, func(p, q int) bool {
        return calorie_sums[p] > calorie_sums[q]
    })
    return calorie_sums[0] + calorie_sums[1] + calorie_sums[2]
}

func main() {
    bytes, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}
