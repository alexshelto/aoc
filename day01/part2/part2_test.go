package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`, 45000},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := compute(tp.in)
        if got != tp.out {
            t.Errorf("got: %q, wanted: %q", got, tp.out)
        }
    }
}
