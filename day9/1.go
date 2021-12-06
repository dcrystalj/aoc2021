package day9

import (
	"math"
	"sort"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)


func init() {
    registrator.Register("day9.1", Solve)
    registrator.Register("day9.2", Solve2)
}

func makeGrid(lines *[]string) [][]int {
	grid := make([][]int, 0)
	for row := 0; row < len((*lines)); row++ {
		gridRow := make([]int, 0)
		for col := 0; col < len((*lines)[row]); col++ {
			val := int((*lines)[row][col]-'0')
			gridRow = append(gridRow, val)
		}

		grid = append(grid, gridRow)
	}
	return grid
}

type Loc struct {
	row, col int
}

func getMinimumsAndRisk(grid *[][]int) (minimums []Loc, risk int){
	n := len((*grid))
	m := len((*grid)[0])
	xdirections := []int{1, -1, 0, 0}
	ydirections := []int{0, 0, 1, -1}
	minimums = make([]Loc, 0)
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			minimumNeighbor := math.MaxInt
			for i := 0; i < 4; i++ {
				newx := col + xdirections[i]
				newy := row + ydirections[i]
				if 0 <= newx && newx < m && 0 <= newy && newy < n {
					minimumNeighbor = utils.MinInt(minimumNeighbor, (*grid)[newy][newx])
				}
			}
			if minimumNeighbor > (*grid)[row][col] {
				risk += (*grid)[row][col] + 1
				minimums = append(minimums, Loc{row, col})
			}
		}
	}
	return minimums, risk
}

func Solve(lines []string) string {
	grid := makeGrid(&lines)
	_, risk := getMinimumsAndRisk(&grid)
	return strconv.Itoa(risk)
}

func Solve2(lines []string) string {
	grid := makeGrid(&lines)
	minimums, _ := getMinimumsAndRisk(&grid)
	basinSizes := make([]int, 0)
	for _, loc := range minimums {
		basinSizes = append(basinSizes, bfsSize(&grid, loc))
	}
	sort.SliceStable(basinSizes, func(a, b int) bool {
		return basinSizes[a] > basinSizes[b]
	})
	return strconv.Itoa(basinSizes[0] * basinSizes[1] * basinSizes[2])
}

type Queue struct {
	loc Loc
	depth int
}

func bfsSize(grid *[][]int, loc Loc) int {
	visited := make(map[Loc]bool, 0)
	d := deque.New()
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d.PushLeft(Queue{loc, 1})
	cnt := 0
	for !d.Empty() {
		head := d.PopLeft().(Queue)
		if _, exists := visited[head.loc]; exists {
			continue
		}
		cnt += 1
		visited[head.loc] = true

		for _, direction := range directions {
			newx := head.loc.col + direction[0]
			newy := head.loc.row + direction[1]
			if 0 <= newx && newx < len((*grid)[0]) && 0 <= newy && newy < len((*grid)) && (*grid)[newy][newx] != 9 {
				d.PushRight(Queue{Loc{newy, newx}, head.depth + 1})
			}
		}
	}
	return cnt
}
