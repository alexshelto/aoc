package main 

import "testing"

var expected = "MCD"

func TestSolution(t *testing.T) {
        got := compute(testInput)
        if got != expected{
            t.Errorf("got: %v, wanted: %v", got, expected)
        }
}
