package main 

import "testing"

var testParams = []struct {
    in []string 
    out int 
} {
    {[]string{"467..114..",
    "...*......",
    "..35..633.",
    "......#...",
    "617*......",
    ".....+.58.",
    "..592.....",
    "......755.",
    "...$.*....",
    ".664.598.."}, 4361},
}


func TestSolution(t *testing.T) {
    for _,tp := range testParams {
        got := compute(tp.in)
        if got != tp.out {
            t.Errorf("got: %q, wanted: %q", got, tp.out)
        }
    }
}


/*
func TestIsAdjacent(t *testing.T) {
    for _, tp := range testParamsIdz {

    }
}
*/
