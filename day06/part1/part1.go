package main

import (
	"fmt"
	"os"
)

var BUFFERSIZE = 4

type Receiver struct {
    buffer []rune
}

func (r *Receiver) Insert(idx int, val rune) {
    r.buffer[idx % BUFFERSIZE] = val
}


func (r *Receiver) AllUnique() bool {
    m := map[rune]bool{}

    for _, val := range r.buffer {
        m[val] = true
    }

    if len(m) != BUFFERSIZE {
        return false
    } 
    return true
}


func compute(input string) int {
    var buf = make([]rune, BUFFERSIZE)
    r := Receiver{buffer: buf}

    for i, c := range input {
        r.Insert(i, c)
        if i >= BUFFERSIZE-1  && r.AllUnique() == true {
            return i + 1
        }
    }
    return -1
}

func main() {
    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}

