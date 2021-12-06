package day11

import (
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
)


func init() {
    registrator.Register("day11.1", Solve)
    registrator.Register("day11.2", Solve2)
}
type Loc struct {
	row int
	col int
}
type Grid struct {
	grid [][]int
	nrows int
	ncols int
	flashed []Loc
}

func (g *Grid) increaseByOne() {
	g.flashed = make([]Loc, 0)
	for row := 0; row < g.nrows; row++ {
		for col := 0; col < g.ncols; col++ {
			g.grid[row][col] += 1
			if g.grid[row][col] > 9 {
				g.flashed = append(g.flashed, Loc{row, col})
			}
		}
	}
}

func (g *Grid) flash() int {
	exploded := make(map[Loc]bool,0)
	for _, loc := range g.flashed {
		exploded[loc] = true
	}
	directions := []Loc{{0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {-1, 0}}
	for len(g.flashed) > 0 {
		loc := g.flashed[len(g.flashed)-1]
		g.flashed = g.flashed[:len(g.flashed)-1]
		for _, n := range directions {
			newx, newy := loc.col + n.col, loc.row + n.row
			newLoc := Loc{newy, newx}
			if 0 <= newx && newx < g.ncols && 0 <= newy && newy < g.nrows {
				g.grid[newy][newx] += 1
				if g.grid[newy][newx] > 9 {
					_, exist := exploded[newLoc]
					if !exist {
						g.flashed = append(g.flashed, newLoc)
						exploded[newLoc] = true
					}
				}
			}
		}
		exploded[loc] = true
	}

	for loc := range exploded {
		g.grid[loc.row][loc.col] = 0
	}
	return len(exploded)
}

func NewGrid(lines []string) *Grid {
	g := Grid{}
	g.nrows, g.ncols = len(lines), len(lines[0])
	g.grid = make([][]int, len(lines))
	for row, line := range lines {
		g.grid[row] = make([]int, len(line))
		for col, val := range line {
			g.grid[row][col] = int(val-'0')
		}
	}
	return &g
}

func Solve(lines []string) string {
	g := NewGrid(lines)
	cnt := 0
	for i := 0; i < 100; i++ {
		g.increaseByOne()
		cnt += g.flash()

	}
	return strconv.Itoa(cnt)
}

func Solve2(lines []string) string {
	g := NewGrid(lines)
	for i := 0; i < 100000; i++ {
		g.increaseByOne()
		cnt := g.flash()
		if cnt == g.ncols * g.nrows {
			return strconv.Itoa(i+1)
		}

	}
	return "-1"
}
