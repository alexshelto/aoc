package main 

import "testing"

var testParams = []struct {
    in string 
    out string 
} {
    {"two1nine", "29"},
    {"eightwothree", "83"},
    {"abcone2threexyz", "13"},
    {"xtwone3four", "24"},
    {"zoneight234", "14"},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := firstAndLastNumber(tp.in)
        if got != tp.out {
            t.Errorf("got: %q, wanted: %q", got, tp.out)
        }
    }
}
