package main

import (
	"fmt"
	"os"
	"strings"
)

type Grid struct {
    numRows int 
    numCols int 
    grid [][]int
}

func NewGrid(input string) Grid {
    lines := strings.Split(input, "\n") 
    numRows := len(lines)
    numCols := len(lines[0])

    fmt.Println("grid is", numRows, "by", numCols)

    vec := make([][]int, numRows)
    for i:=0; i < numRows; i++ {
        vec[i] = make([]int, numCols)
    }

    for rowNum, line := range lines {
        for idx, char := range line {
            val := int(char - '0')
            vec[rowNum][idx] = val
        }
    }
    return Grid{numRows: numRows, numCols: numCols, grid: vec}
} 


func (g *Grid) IsEdge(r int, c int) bool {
    if r == 0 || r == g.numRows - 1 || c == 0 || c == g.numCols - 1 {
        fmt.Print("is Edge")
        return true

    }
    return false
}



func (g *Grid) VisibleUp(r int, c int) bool {
    // Moving row idx [row][col] by -1
    curVal := g.grid[r][c]
    for rowIdx := r-1; rowIdx >= 0; rowIdx-- {
        if g.grid[rowIdx][c] >= curVal {
            return false
        }
    }
    fmt.Print("VisibleUp ")
    return true
}

func (g *Grid) VisibleDown(r int, c int) bool {
    curVal := g.grid[r][c]
    for rowIdx := r+1; rowIdx < g.numRows; rowIdx++ {
        if g.grid[rowIdx][c] >= curVal {
            return false
        }
    }
    fmt.Print("VisibleDown ")
    return true
}


func (g *Grid) VisibleLeft(r int, c int) bool {
    curVal := g.grid[r][c]
    for colIdx := c-1; colIdx >= 0; colIdx-- {
        if g.grid[r][colIdx] >= curVal {
            return false
        }
    }
    fmt.Print("VisibleLeft ")
    return true
}


func (g *Grid) VisibleRight(r int, c int) bool {
    curVal := g.grid[r][c]
    for colIdx := c+1; colIdx < g.numCols; colIdx++ {
        if g.grid[r][colIdx] >= curVal {
            return false
        }
    }
    fmt.Print("VisibleRight ")
    return true
}


func (g *Grid) LocationIsVisible(r int, c int) bool {
    if g.VisibleDown(r,c) || g.VisibleUp(r,c) || g.VisibleLeft(r,c) || g.VisibleRight(r,c) {
        return true
    }
    return false
}


func compute(input string) int {
    grid := NewGrid(input)
    ret := 0

    for _, row := range grid.grid {
        fmt.Println(row)
    }

    for i:=0; i < grid.numRows; i++ {
        for j:=0; j < grid.numCols; j++ {
            if grid.IsEdge(i, j) || grid.LocationIsVisible(i, j) {
                fmt.Println("Value:", grid.grid[i][j], "row", i, "col", j)
                ret += 1
            }
        }
        fmt.Println("##")
    }
    fmt.Println(grid)
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

