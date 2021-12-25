package day20

import (
	"fmt"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
)

func init() {
	registrator.Register("day20.1", Solve)
	registrator.Register("day20.2", Solve2)

}

func Solve(lines []string) string {
	return solveForStep(lines, 2)
}

func Solve2(lines []string) string {
	return solveForStep(lines, 50)
}

func solveForStep(lines []string, steps int) string {
	enhancmentMap := createEnhancmentMap(lines[0])
	nrows := len(lines) - 2
	ncols := len(lines[3])
	m := NewMap(lines[2:])
	start, end := -steps*2, len(lines)+steps*2
	for i := 0; i < steps; i++ {
		m = transformMap(m, start, end, enhancmentMap)
		start -= 2
		end += 2
	}
	// printMap(0-steps, 0-steps, start, end, m)
	return strconv.Itoa(countLights(m, 0-steps, 0-steps, nrows+steps, ncols+steps))
}

type Point struct {
	y, x int
}

type Map map[Point]bool

func NewMap(lines []string) *Map {
	m := make(Map)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == '#' {
				m[Point{i, j}] = true
			}
		}
	}
	return &m
}

func createEnhancmentMap(s string) []bool {
	b := make([]bool, len(s))
	for i, c := range s {
		if c == '#' {
			b[i] = true
		}
	}
	return b
}

func countLights(m *Map, startrow, startcol, endrow, endcol int) (cnt int) {
	for i := startrow; i < endrow; i++ {
		for j := startcol; j < endcol; j++ {
			_, isLid := (*m)[Point{i, j}]
			if isLid {
				cnt += 1
			}
		}
	}
	return
}

func transformMap(m *Map, start int, end int, enhancmentMap []bool) *Map {
	m1 := make(Map)
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			if enhancePoint(m, Point{i, j}, enhancmentMap) {
				m1[Point{i, j}] = true
			}
		}
	}
	return &m1
}

func enhancePoint(m *Map, p Point, enhancmentMap []bool) bool {
	binStr := ""
	for i := p.y - 1; i <= p.y+1; i++ {
		for j := p.x - 1; j <= p.x+1; j++ {
			_, isLight := (*m)[Point{i, j}]
			if isLight {
				binStr += "1"
			} else {
				binStr += "0"
			}
		}
	}
	bin, _ := strconv.ParseInt(binStr, 2, 32)
	return enhancmentMap[bin]
}

// should account for flipping background here
// func enhanceForPoint(m *Map, p Point, enhancmentMap []bool) []Point {
// 	enhanced := make([]Point, 0)
// 	for a := 0; a < 3; a++ {
// 		for b := 0; b < 3; b++ {
// 			binStr := ""
// 			for i := p.y - 2 + a; i <= p.y+a; i++ {
// 				for j := p.x - 2 + b; j <= p.x+b; j++ {
// 					_, isLight := (*m)[Point{i, j}]
// 					if isLight {
// 						binStr += "1"
// 					} else {
// 						binStr += "0"
// 					}
// 				}
// 			}
// 			bin, _ := strconv.ParseInt(binStr, 2, 32)
// 			if enhancmentMap[bin] {
// 				enhanced = append(enhanced, Point{p.y - 1 + a, p.x - 1 + b})
// 			}
// 		}
// 	}
// 	return enhanced
// }

func printMap(startrow int, startcol int, start int, end int, m *Map) {
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			if i == startrow && j == startcol {
				fmt.Print("A")
				continue
			}
			_, found := (*m)[Point{i, j}]
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
