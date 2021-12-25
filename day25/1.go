package day25

import (
	"fmt"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
)

func init() {
	registrator.Register("day25.1", Solve)
	registrator.Register("day25.2", Solve2)

}

func Solve(lines []string) string {
	cm := mapGrid(lines)
	r := cm.countMoves()
	return strconv.Itoa(r)
}

func Solve2(lines []string) string {
	return strconv.Itoa(3)
}

type Dir int

const (
	NONE Dir = iota
	R
	D
)

type Dim struct {
	rows, cols int
}

type Loc struct {
	row, col int
}

// type Cucumber struct {
// 	row, col int
// 	toRight bool
// }

type CucumberMap struct {
	m   map[Loc]Dir
	dim Dim
}

func NewDim(lines []string) Dim {
	return Dim{len(lines), len(lines[0])}
}

func mapGrid(lines []string) CucumberMap {
	dim := NewDim(lines)
	cm := make(map[Loc]Dir)
	for row, line := range lines {
		for col, c := range line {
			if c == '>' {
				cm[Loc{row, col}] = R
			} else if c == 'v' {
				cm[Loc{row, col}] = D
			}
		}
	}
	return CucumberMap{cm, dim}
}

func (cm *CucumberMap) countMoves() (cnt int) {
	for (*cm).move() {
		cnt += 1
	}
	return cnt + 1
}

func (cm *CucumberMap) move() bool {
	tmpMap, hasMoved1 := cm.moveRight()
	hasMoved2 := cm.moveDown(tmpMap)
	// cm.print()
	return hasMoved1 || hasMoved2
}

func (cm *CucumberMap) moveDown(mcm *CucumberMap) (hasMoved bool) {
	cmP := *cm
	for loc, dir := range cmP.m {
		if dir == D {
			newLoc := Loc{(loc.row + 1) % cm.dim.rows, loc.col}
			newLocCurrentVal, found := cmP.m[newLoc]
			newLocNextVal, found1 := mcm.m[newLoc]
			if found && newLocCurrentVal == D || found1 && newLocNextVal == R {
				mcm.m[loc] = D
			} else {
				hasMoved = true
				mcm.m[newLoc] = D
			}
		}
	}
	*cm = *mcm
	return
}

func (cm *CucumberMap) moveRight() (*CucumberMap, bool) {
	hasMoved := false
	cmP := *cm
	mcm := &CucumberMap{make(map[Loc]Dir), cmP.dim}
	for loc, dir := range cmP.m {
		if dir == R {
			newLoc := Loc{loc.row, (loc.col + 1) % cm.dim.cols}
			if cmP.isOccupied(newLoc) {
				mcm.m[loc] = R
			} else {
				hasMoved = true
				mcm.m[newLoc] = R
			}
		}
	}
	return mcm, hasMoved
}

func (cm *CucumberMap) isOccupied(loc Loc) bool {
	_, found := (*cm).m[loc]
	return found
}

func (cm *CucumberMap) print() {
	cmP := *cm
	for i := 0; i < cmP.dim.rows; i++ {
		for j := 0; j < cmP.dim.cols; j++ {
			dir, found := cmP.m[Loc{i, j}]
			if found {
				fmt.Print(dir)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
