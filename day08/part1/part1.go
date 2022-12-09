package main

import (
	"fmt"
	"os"
	"strings"
    "github.com/alexshelto/aoc2022/utils"
)



func Max(a []uint8) uint8 {
    ret := a[0]
    for _, v := range a[1:] {
        if v > ret {
            ret = v
        }
    }
    return ret
}



func compute(input string) int {
    input = strings.TrimSuffix(input, "\n")
    lines := strings.Split(input, "\n")

    matrix := utils.BuildDigitMatrix(lines)


    // going left
    //fmt.Println(matrix[2][:2])

    n := len(matrix)
    m := len(matrix[0])

    ret := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            val := matrix[i][j]

            // L
            if j == 0 || Max(matrix[i][:j]) < val {
                ret += 1
                continue
            }
            //R
            if j == m - 1 || Max(matrix[i][(j+1):]) < val {
                ret += 1
                continue
            }
            // up 
            if i == 0 || Max(matrix.ColumnAtIndex(j)[:i]) < val {
                ret += 1
                continue
            }
            // down 
            if i == n - 1 || Max(matrix.ColumnAtIndex(j)[(i+1):]) < val {
                ret += 1
                continue
            }
        }
    }

    return ret
}


func main () {

    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))

}
