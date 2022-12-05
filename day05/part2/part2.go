package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
    "errors"
	"strings"
)


var testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

var EmptyStackOnPopError = "Empty Stack, Cant Pop"

type Stack struct {
    contents []string 
}


func (s *Stack) Push(v string) {
    s.contents = append([]string{v}, s.contents...)
}

// make pop
func (s *Stack) Pop() (string, error) {
    length := len(s.contents)
    if length < 1 {
        return "", errors.New(EmptyStackOnPopError)
    }
    val := s.contents[length - 1]
    s.contents = s.contents[0:length-1]
    return val, nil
}



type CraneYard struct {
    stacks []Stack
}

func New(vec []string) CraneYard {
    n_lines := len(vec)
    cy := CraneYard{}

    for i := 0; i < n_lines + 1; i++ {
        cy.stacks = append(cy.stacks, Stack{})
    }

    for _, line := range vec {
        for i, c := range line {
            if unicode.IsLetter(c) {
                stackNo := (int(vec[n_lines-1][i]) - '0')
                cy.stacks[stackNo].Push(string(c))
            }
        }
    }
    return cy
}


func compute(input string) string {
    lines := strings.Split(input, "\n\n")
    setupChunk := strings.Split(lines[0], "\n")
    moves := strings.Split(lines[1], "\n")

    cy := New(setupChunk)
    fmt.Println(cy)

    re := regexp.MustCompile("([0-9]+)") 

    for _, line := range moves {
        if line == "" {
            break
        }
        matches := re.FindAllString(line, -1)

        count, _ := strconv.Atoi(matches[0]) 
        from, _ := strconv.Atoi(matches[1]) 
        to, _ := strconv.Atoi(matches[2]) 

        tmp_stack := Stack{}

        for i := 0; i < count; i++ {
            val, err := cy.stacks[from].Pop()
            if err != nil {
                panic("")
            }
            tmp_stack.contents = append(tmp_stack.contents, val)
        }

        for i := 0; i < count; i++ {
            val, err := tmp_stack.Pop()
            if err == nil {
                cy.stacks[to].contents = append(cy.stacks[to].contents, val)
            }
        }
    }


    ret := "" 

    for _, stack := range cy.stacks {
        val, err := stack.Pop()
        if err == nil {
            ret = ret + val
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

