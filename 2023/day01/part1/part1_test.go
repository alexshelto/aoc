package main 

import "testing"

var testParams = []struct {
    in string 
    out string 
} {
    {"1abc2", "12"},
    {"pqr3stu8vwx", "38"},
    {"a1b2c3d4e5f", "15"},
    {"treb7uchet", "77"},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := firstAndLastNumber(tp.in)
        if got != tp.out {
            t.Errorf("got: %q, wanted: %q", got, tp.out)
        }
    }
}
