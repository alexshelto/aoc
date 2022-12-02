package main 

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

var remap = map[string]string{"A":"R", "B":"P", "C":"S"}
var winsTo = map[string]string{"R":"S", "S":"P", "P":"R"}
var loosesTo = map[string]string{"R":"P", "P":"S", "S": "R"}
var scoreTable = map[string]int{"R":1, "P":2, "S":3}

func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    ret := 0

    for scanner.Scan() {
        moves := strings.Split(scanner.Text(), " ")
        theirHand := remap[moves[0]]
        myHand := moves[1]

        if myHand == "X" {
            ret += scoreTable[winsTo[theirHand]]
        } else if myHand == "Y" {
            ret += scoreTable[theirHand] + 3
        } else {
            ret += scoreTable[loosesTo[theirHand]] + 6
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
