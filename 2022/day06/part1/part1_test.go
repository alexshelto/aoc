package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
    {"nppdvjthqldpwncqszvftbrmjlhg", 6},
    {"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
    {"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
}



func TestSolution(t *testing.T) {
    for _,tp := range testParams {
            got := compute(tp.in)
            if got != tp.out{
                t.Errorf("got: %v, wanted: %v", got, tp.out)
            }
        }
}
