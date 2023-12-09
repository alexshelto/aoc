package main 

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func max(num1 int, num2 int) int {
    if num1 >= num2 {
        return num1 
    }
    return num2
}


func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))

    ret := 0
    local_sum := 0

    for scanner.Scan() {
        line := scanner.Text()
        // If we hit newline new group 
        if len(line) == 0 {
            ret = max(local_sum, ret) 
            local_sum = 0
        } else {
            calories, err := strconv.Atoi(line)
            if err != nil {
                panic("attempted to parse invalid int")
            }
            local_sum += calories 
        }
    }
    return ret
}

func main() {
    bytes, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}
