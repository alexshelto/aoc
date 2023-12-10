package main

import "testing"

var testParams = []struct {
	in  []string
	out int
}{
	{[]string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}, 467835},
}

func TestSolution(t *testing.T) {
	for _, tp := range testParams {
		got := compute(tp.in)
		if got != tp.out {
			t.Errorf("got: %q, wanted: %q", got, tp.out)
		}
	}
}

func TestIsAdjacent(t *testing.T) {
    p1 := Point{2,3}
    p2 := Point{1,3}

    got := arePointsAdjacent(p1, p2)

    if got != true {
        t.Error("Got: ", got, " wanted: ", true)
    }
}
