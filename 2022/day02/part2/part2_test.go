package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {`A Y
B X
C Z`, 12},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := compute(tp.in)
        if got != tp.out {
            t.Errorf("got: %v, wanted: %v", got, tp.out)
        }
    }
}
