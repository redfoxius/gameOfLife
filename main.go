package main

import (
	"fmt"
	"time"
)

const (
	Size = 25
)

const (
	Dead = iota
	Alive
)

type (
	GridUniverse [Size][Size]uint8 // what will be here with Size=1000000 for example?
	Universe     interface {
		Tick() Universe
		Print()
	}
)

func (gu GridUniverse) countCellAliveNeighbors(x, y int) uint16 {
	var count uint16
	/*
	Really strange choose of type. If you are using uint, then looks like you are trying to optimize memory. But the problem is 
 	next: your max value will be 8, so its enough uint8 to keep it.
 	*/
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < Size && ny >= 0 && ny < Size && gu[ny][nx] == Alive {
				count++
			}
		}
	}

	return count
}

func (gu GridUniverse) Tick() Universe {
	// this function is useless:
	nextGeneration := seedGridUniverse()
	// You can just use next line. Result is the same, but without context switch and useless full slice walk
	// var universe GridUniverse

	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			liveNeighbors := gu.countCellAliveNeighbors(x, y)

			if gu[y][x] == Alive {
				switch {
				case liveNeighbors < 2:
					nextGeneration[y][x] = Dead
				case liveNeighbors > 3:
					nextGeneration[y][x] = Dead
				default:
					nextGeneration[y][x] = Alive
				}
			} else {
				switch {
				case liveNeighbors == 3:
					nextGeneration[y][x] = Alive
				default:
					nextGeneration[y][x] = Dead
				}
			}
		}
	}

	return nextGeneration
}

// Why this var is located here? Between functions. Why we can't move it to the top?
var glider = [][]uint8{
	{Dead, Alive, Dead},
	{Dead, Dead, Alive},
	{Alive, Alive, Alive},
}

func seedGridGliderUniverse() Universe {
	universe := seedGridUniverse()

	startX := (Size - len(glider)) / 2
	startY := (Size - len(glider[0])) / 2

	for y, row := range glider {
		for x, val := range row {
			universe[startY+y][startX+x] = val
		}
	}

	return universe
}

// as i wrote before, this function is totally useless and will slow application
func seedGridUniverse() GridUniverse {
	var universe GridUniverse
	for y := 0; y < Size; y++ {
		// Next construction will be much faster:
		// universe[y] = make([]uint8, 0, Size)
		for x := 0; x < Size; x++ {
			universe[y][x] = Dead
		}
	}

	return universe
}

func (gu GridUniverse) Print() {
	for _, row := range gu {
		for _, cell := range row {
			if cell == Alive {
				fmt.Print("* ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	universe := seedGridGliderUniverse()

	// why we need infinite loop if nothing is changed after 47 iterations?
	for i := 0; ; i++ {
		fmt.Printf("Generation %d:\n", i)
		universe.Print()
		universe = universe.Tick()
		time.Sleep(200 * time.Millisecond)
	}
}
