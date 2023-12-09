package main 

import (
    "fmt"
    "os"
    "strings"
)

/*
-------------------
read file as string 
-------------------
bytes, err := os.ReadFile(name)
str := string(bytes)

-----------------------
read file line by line 
-----------------------
readFile, err := os.Open(name)
fileScanner := bufio.NewScanner(readFile)
for fileScanner.Scan() {}
*/

func compute(input string) int {
    ret := 0
    for _, c := range strings.TrimSuffix(input, "\n") {
    }
    return ret
}

func main() {
    bytes, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    input := string(bytes)
    fmt.Println(compute(input))
}
