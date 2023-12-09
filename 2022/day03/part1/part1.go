package main 

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

/*
lowercase item types a - z have priorities 1 - 26
uppercase item types A - Z have priorities 27 - 52

ASCII: A = 65 ... ASCII - 38 if uppercase 
ASCII: a = 96 ... ASCII - 96 if lowercase 
*/
func PrioValue(asciiValue int) int {
    if asciiValue >= 65 && asciiValue <= 90 {
        return asciiValue - 38
    } else {
        return asciiValue - 96
    }
}



func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    ret := 0

    for scanner.Scan() {
        line := scanner.Text()
        length := len(line)

        set := map[rune]bool{}
        compartment1 := line[0:length/2]
        compartment2 := line[length/2:]

        for _, c := range compartment1 {
            set[c] = true
        }

        for _, c := range compartment2 {
            if set[c] == true {
                ret += PrioValue(int(c))
                delete(set, c) 
            }
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
