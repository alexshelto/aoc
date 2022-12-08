package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

var MAX_DIR_SIZE = 100000



type Set map[string]bool

func (s *Set) Add(val string) {
    (*s)[val] = true
}


func compute(input string) int {

    lines := strings.Split(input, "\n")
    dirs := Set{"/": true}
    files := map[string]int{}
    pwd := "/"

    for _, line := range lines[1:] {
        wds := strings.Split(line, " ")

        fmt.Println(line)

        if strings.HasPrefix(line, "$ cd ..") {
            if pwd == "/" {
            } else {
                pwd = filepath.Dir(pwd)
            }
        } else if strings.HasPrefix(line, "$ cd"){
            pwd = filepath.Join(pwd, wds[2])
        } else if strings.HasPrefix(line, "$ ls") {
        } else if strings.HasPrefix(line, "dir") {
            dirs.Add(filepath.Join(pwd, wds[1]))
        } else if line == "" {
            continue
        } else {
            size, _ := strconv.Atoi(wds[0])
            files[filepath.Join(pwd, wds[1])] = size
        }
    }

    root := 0
    for _, v := range files {
        root += v
    }
    if root > MAX_DIR_SIZE {
        root = 0
    }

    for dname, _ := range dirs {
        root += Size(dname, files)
    }

    return root
}


func Size(dirname string, m map[string]int) int {
    size := 0
    for k, v := range m {
        if strings.HasPrefix(k, dirname+"/") {
            size += v
        }
    }
    if size > MAX_DIR_SIZE {
        return 0
    } else {
        return size
    }
}


func main() {
    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}

