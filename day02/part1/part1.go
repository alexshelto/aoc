package main 

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

var loosesTo = map[string]string{"A":"Y", "B":"Z", "C": "X"}
var tiesTo = map[string]string{"A":"X", "B":"Y", "C": "Z"}

func calcMoveScore(throw string) int {
    if throw == "X" {
        return 1
    } else if throw == "Y" {
        return 2
    } else {
        return 3
    }
}


func compute(input string) int {
    scanner := bufio.NewScanner(strings.NewReader(input))
    ret := 0

    for scanner.Scan() {
        line := scanner.Text()
        moves := strings.Split(line, " ")

        theirHand := moves[0]
        myHand := moves[1]

        if loosesTo[theirHand] == myHand {
            ret += calcMoveScore(myHand) + 6
        } else if tiesTo[theirHand] == myHand {
            ret += calcMoveScore(myHand) + 3
        } else {
            ret += calcMoveScore(myHand)
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
