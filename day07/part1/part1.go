package main

import (
	"fmt"
	"os"
    "strings"
    "strconv"
)

var MAX_SIZE_SUM = 100000


type DirectoryChildren map[string][]string 
func (dc *DirectoryChildren) AddDescendent(path string, dir string) {
    if _, ok := (*dc)[path]; ok {
        (*dc)[path] = append((*dc)[path], dir)
    } else {
        (*dc)[path] = []string{dir}
    }
}


type DirectorySizeTable map[string]int

func (dst *DirectorySizeTable) IncrementBy(path string, size int) {
    if val, ok := (*dst)[path]; ok {
        (*dst)[path] = val + size
    } else {
        (*dst)[path] = val
    }
}




type PathStack struct {
    paths []string 
}

func (p *PathStack) Push(fp string) {
    p.paths = append(p.paths, fp)

    fmt.Println("Pushed to stack")
}

func (p *PathStack) Pop() string {
    length := len(p.paths)

    if length < 1 {
        panic("Empty stack, this is wrong")
    }

    val := p.paths[length-1]
    p.paths = p.paths[0:length-1] 

    return val
}

func (p * PathStack) CurrentPath() string {
    length := len(p.paths)

    if length < 1 {
        panic("Empty stack, this is wrong")
    }

    val := p.paths[length-1]
    return val
}



func compute(input string) int {
    lines := strings.Split(input, "\n")

    pathStack := PathStack{}
    dirChildren := DirectoryChildren{}
    sizeTable := DirectorySizeTable{}

    for _, line := range lines {

        if line[0] == '$' {
            listingOutput = false
        }


        if listingOutput {
            wds := strings.Split(line, " ")
            if wds[0] == "dir" {
                dirChildren.AddDescendent(pathStack.CurrentPath(), wds[1])
            } else {
                size, err := strconv.Atoi(wds[0])
                if err != nil {
                    panic("Parsed wrong")
                }
                sizeTable.IncrementBy(pathStack.CurrentPath(), size)
            }
        }
        if line[0] == '$' {
            listingOutput = false
            if line == "$ ls" {
                listingOutput = true
            } else {
                wds := strings.Split(line, " ") 
                if wds[1] == ".." {
                    pathStack.Pop()
                } else {
                    pathStack.Push(wds[2])
                }
            }
        }


        fmt.Println("input:", line, " | stack:", pathStack.paths, " sizes:", sizeTable)

    }

    fmt.Println(sizeTable)
    fmt.Println(dirChildren)

    return 0
}

func main() {
    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}

