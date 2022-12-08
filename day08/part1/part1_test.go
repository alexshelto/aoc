package main 

import (
    "testing"
)
var testInput = `30373
25512
65332
33549
35390`


var testInput2 = `31373
20002
00102
30009
31117`


var expectedResult = 21


func TestSolution(t *testing.T) {
    got := compute(testInput)
    if got != expectedResult{
        t.Errorf("got: %v, wanted: %v", got, expectedResult)
    }
}


func TestLeft(t *testing.T) {
    grid := NewGrid(testInput2)
    my := grid.VisibleLeft(2,2)
    expected := true
    if my != expected {
        t.Errorf("got: %v, wanted: %v", my, expected)
    }

}

func TestUp(t *testing.T) {
    grid := NewGrid(testInput2)
    my := grid.VisibleDown(2,3)
    expected := false
    if my != expected {
        t.Errorf("got: %v, wanted: %v", my, expected)
    }

}


func TestDown(t *testing.T) {
    grid := NewGrid(testInput2)
    my := grid.VisibleDown(2,2)
    expected := false
    if my != expected {
        t.Errorf("got: %v, wanted: %v", my, expected)
    }

}
