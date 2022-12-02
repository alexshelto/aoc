package main 

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

var remap = map[string]string{"A":"R", "B":"P", "C":"S",
                              "X":"R", "Y":"P", "Z":"S",
                             }
var winsTo = map[string]string{"R":"S", "S":"P", "P":"R"}
var loosesTo = map[string]string{"R":"P", "P":"S", "S": "R"}
var scoreTable = map[string]int{"R":1, "P":2, "S":3}


func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    ret := 0

    for scanner.Scan() {
        moves := strings.Split(scanner.Text(), " ")
        theirHand := remap[moves[0]]
        myHand := remap[moves[1]]

        if loosesTo[theirHand] == myHand {
            ret += scoreTable[myHand] + 6
        } else if theirHand == myHand {
            ret += scoreTable[myHand] + 3
        } else {
            ret += scoreTable[myHand]
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
