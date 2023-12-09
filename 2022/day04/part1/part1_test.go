package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`, 2},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := compute(tp.in)
        if got != tp.out {
            t.Errorf("got: %v, wanted: %v", got, tp.out)
        }
    }
}
