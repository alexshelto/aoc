package main 

import (
    "testing"
)
var testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`


var expectedResult = 13

func TestSolution(t *testing.T) {
    got := compute(testInput)
    if got != expectedResult{
        t.Errorf("got: %v, wanted: %v", got, expectedResult)
    }
}

