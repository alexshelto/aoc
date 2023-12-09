package main 

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func PrioValue(asciiValue int) int {
    if asciiValue >= 65 && asciiValue <= 90 {
        return asciiValue - 38
    } else {
        return asciiValue - 96
    }
}


type SharedCharSet map[rune]int

func (scs *SharedCharSet) RecordChars(ruck string) {
    history := map[rune]bool{}

    for _, c := range ruck {
        _, alreadyCounted := history[c]

        if !alreadyCounted {
            _, ok := (*scs)[c]
            if !ok {
                (*scs)[c] = 1
            } else {
                (*scs)[c] += 1
            }
            history[c] = true
        }
    }
}

func (scs *SharedCharSet) RecordedThreeTimes() rune {
    for k, v := range (*scs) {
        if v == 3 {
            return k
        }
    }
    panic("nothing was recorded 3 times")
}



func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    ret := 0

    var lines []string 
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    for i := 0; i < len(lines); i += 3 {
        s := &SharedCharSet{}
        s.RecordChars(lines[i])
        s.RecordChars(lines[i + 1])
        s.RecordChars(lines[i + 2])

        ret += PrioValue(int(s.RecordedThreeTimes()))
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
