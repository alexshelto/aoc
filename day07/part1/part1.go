package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MAX_DIR_SIZE = 100000

type PathStack struct {
    paths []string 
}

func (ps *PathStack) ChangeDir(path string) {
    if path == ".." {
        ps.Pop()
    } else {
        ps.Push(path)
    }
}

func (ps *PathStack) Push(path string) {
    ps.paths = append(ps.paths, path)
}

func (ps *PathStack) CurrentPath() string {
    length := len(ps.paths)

    if length < 1 {
        panic("Blew up in curentPath() something went wrong parsing")
    }
    val := ps.paths[length - 1]
    return val
}

func (ps *PathStack) Pop() string {
    length := len(ps.paths)

    if length < 1 {
        panic("Blew up in Pop(), something went wrong parsing")
    }
    val := ps.paths[length - 1]
    ps.paths = ps.paths[0:length-1]

    return val
}

type DirectoryTable map[string][]string

func (dt *DirectoryTable) AddTo(dirName string, path string) {
    if _, ok := (*dt)[path]; ok {
        (*dt)[path] = append((*dt)[path], dirName) 
    } else {
        (*dt)[path] = []string{dirName}
    }
}

type DirSize map[string]int

func (ds *DirSize) AddTo(size int, path string) {
    if current, ok := (*ds)[path]; ok {
        (*ds)[path] = current + size 
    } else {
        (*ds)[path] = size 
    }
}



func compute(input string) int {
    lines := strings.Split(input, "\n")

    dirTable := DirectoryTable{}
    pathStack := PathStack{}
    dirSizes := DirSize{}

    for _, line := range lines {
        wds := strings.Split(line, " ")

        if line[0] == '$' {
            fmt.Println("command:", line)
            if wds[1] == "cd" {
                pathStack.ChangeDir(wds[2])
            } else {
            }
        } else {
            fmt.Println("output:", line)
            if wds[0] == "dir" {
                dirTable.AddTo(wds[1], pathStack.CurrentPath())
            } else {
                val, _ := strconv.Atoi(wds[0])
                fmt.Println(val)
                dirSizes.AddTo(val, pathStack.CurrentPath())
            }
        }
    }

    /*
        map[/:23352670 a:94269 d:24933642 e:584]
        map[/:[a d] a:[e]]
    */

    ret := 0

    for dir, nestedDirs := range dirTable {
        fmt.Println("Dir:", dir, "Contains dirs:", nestedDirs)
        val, _ := dirSizes[dir]
        tmp := val
        for _, nestedDir := range nestedDirs {
            tmp += dirSizes[nestedDir]
        }

        if tmp <= MAX_DIR_SIZE {
            ret += tmp
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

