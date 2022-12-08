package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
    "sort"
	"strconv"
)


var TOTAL_DISK_SPACE = 70000000
var MIN_SPACE = 30000000

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


    sizes := []int{}

    root := 0
    for _, v := range files {
        root += v
    }
    sizes = append(sizes, root)
    for dname, _ := range dirs {
        sizes = append(sizes, Size(dname, files))
    }

    sort.Slice(sizes, func(i, j int) bool {
        return sizes[i] < sizes[j]
    })

    threshold := MIN_SPACE - (TOTAL_DISK_SPACE - root)

    acceptableSlices := []int{}

    for _, v := range sizes {
        if v >= threshold {
            acceptableSlices = append(acceptableSlices, v)
        }
    }

    sort.Slice(acceptableSlices, func(i, j int) bool {
        return sizes[i] < sizes[j]
    })

    return acceptableSlices[0]
}



func Size(dirname string, m map[string]int) int {
    size := 0
    for k, v := range m {
        if strings.HasPrefix(k, dirname+"/") {
            size += v
        }
    }
    return size
}


func main() {
    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))
}

