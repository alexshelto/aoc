package main 

import (
    "testing"
)
var testInput = `30373
25512
65332
33549
35390`



var expectedResult = 21


func TestSolution(t *testing.T) {
    got := compute(testInput)
    if got != expectedResult{
        t.Errorf("got: %v, wanted: %v", got, expectedResult)
    }
}

