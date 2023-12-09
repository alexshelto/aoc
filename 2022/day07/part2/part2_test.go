package main 

import "testing"

var testParams = []struct {
    in string 
    out int 
} {
    {"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
    {"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
    {"nppdvjthqldpwncqszvftbrmjlhg", 23},
    {"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
    {"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
}



func TestSolution(t *testing.T) {
    for _,tp := range testParams {
            got := compute(tp.in)
            if got != tp.out{
                t.Errorf("got: %v, wanted: %v", got, tp.out)
            }
        }
}
