package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)




type Coordinate struct {
    x, y int 
}


func Max(n1 int, n2 int) int {
    if n1 >= n2 {
        return n1
    }
    return n2
}


// return max of abs(x-x) and (y-y)
func (c *Coordinate) DistanceFrom(c2 *Coordinate) int {
    return Max(
        int(math.Abs( float64(c2.x - c.x) )),
        int(math.Abs( float64(c2.y - c.y) )),
    )
}

func (c *Coordinate) InRowOrColWith(c2 *Coordinate) bool {
    if c.x == c2.x || c.y == c2.y {
        return true
    }
    return false
}



func (c *Coordinate) UpDownLeftRight(head *Coordinate) Coordinate {
    currentX := c.x
    currentY := c.y

    moves := []Coordinate{
        {currentX+1, currentY}, // right 1
        {currentX-1, currentY}, // left 1
        {currentX, currentY+1}, // up 1
        {currentX, currentY-1}, // down 1
    }
    
    for _, coordinate := range moves {
        dist := coordinate.DistanceFrom(head)
        if dist == 1 {
            return coordinate
        }
    }
    panic("damn couldnt find a move")
}

func (c *Coordinate) Diagnol(head *Coordinate) Coordinate {
    currentX := c.x
    currentY := c.y

    moves := []Coordinate{
        {currentX+1, currentY+1}, // diag top right 
        {currentX+1, currentY-1}, // diag bottom right
        {currentX-1, currentY+1}, // diag top left
        {currentX-1, currentY-1}, // diag bottom left
    }
    
    for _, coordinate := range moves {
        dist := coordinate.DistanceFrom(head)
        if dist == 1 {
            return coordinate
        }
    }
    panic("damn couldnt find a move")
}

type Move struct {
    d string
    n int
}

func NewMove(line string) Move {
    wds := strings.Split(line, " ")
    val, _ := strconv.Atoi(wds[1]) 

    return Move{wds[0], val}
}

type Set map[Coordinate]bool 



type Solution struct {
    tailSet Set 
    headCoords Coordinate
    tailCoords Coordinate
}

func (s *Solution) HandleMove(m *Move) {
    for i := 0; i < m.n; i++ {
        // Each step 
        if m.d == "R" {
            s.headCoords.x += 1
        } else if m.d == "L" {
            s.headCoords.x -= 1
        } else if m.d == "U" {
            s.headCoords.y += 1
        } else {
            s.headCoords.y -= 1
        }


        dist := s.headCoords.DistanceFrom(&s.tailCoords)

        fmt.Println("Moved head to:",s.headCoords)
        fmt.Println("Before potential move, tail is at:", s.tailCoords, "distance is:", dist)

        // First check if theyre same row or col.. 
        // then decide the distance.. itl still no matter what be 2 


        if dist > 1  {
            // In row 
            if s.headCoords.InRowOrColWith(&s.tailCoords) {
                newTail := s.tailCoords.UpDownLeftRight(&s.headCoords)
                s.tailCoords = newTail
                s.tailSet[newTail] = true
                fmt.Println("Moved tail to:", newTail)
            } else {
                newTail := s.tailCoords.Diagnol(&s.headCoords)
                s.tailCoords = newTail
                s.tailSet[newTail] = true
                fmt.Println("Moved tail to:", newTail)
            }
        }
    }
}


func compute(input string) int {
    input = strings.TrimSuffix(input, "\n")
    lines := strings.Split(input, "\n")

    headLocation := Coordinate{0,0}
    tailVisits := Set{}
    tailVisits[headLocation] = true
    solution := Solution{tailVisits, Coordinate{0,0}, Coordinate{0,0}}


    for _, line := range lines {
        m := NewMove(line)
        fmt.Println(m)
        solution.HandleMove(&m)
    }

    return len(solution.tailSet)
}

func main() {
    bytes, err := os.ReadFile("../input.txt")
    if err != nil {
        panic(err)
    }
    str := string(bytes)
    fmt.Println(compute(str))

}
