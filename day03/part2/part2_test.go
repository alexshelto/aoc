package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`, 70},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := compute(tp.in)
        if got != tp.out {
            t.Errorf("got: %v, wanted: %v", got, tp.out)
        }
    }
}
