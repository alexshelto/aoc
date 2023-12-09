package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {"(())", 0},
    {"(((", 3},
    {"(()(()(", 3},
    {"())", -1},
    {"))(", -1},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := compute(tp.in)
        if got != tp.out {
            t.Errorf("got: %q, wanted: %q", got, tp.out)
        }
    }
}
