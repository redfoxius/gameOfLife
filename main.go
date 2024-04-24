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
	GridUniverse [Size][Size]uint8
	Universe     interface {
		Tick() Universe
		Print()
	}
)

func (gu GridUniverse) countCellAliveNeighbors(x, y int) uint16 {
	var count uint16
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
	nextGeneration := seedGridUniverse()

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

func seedGridUniverse() GridUniverse {
	var universe GridUniverse
	for y := 0; y < Size; y++ {
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

	for i := 0; ; i++ {
		fmt.Printf("Generation %d:\n", i)
		universe.Print()
		universe = universe.Tick()
		time.Sleep(200 * time.Millisecond)
	}
}
