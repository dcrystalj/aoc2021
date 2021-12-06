package day15

import (
	"fmt"
	"math"
	"strconv"

	"container/heap"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
	"github.com/yourbasic/graph"
)

func init() {
	registrator.Register("day15.1", Solve)
	registrator.Register("day15.2", Solve2)
}

type Grid [][]int

func Solve(lines []string) string {
	nrows := len(lines)
	ncols := len(lines[0])
	g := createGrid(lines)
	dp := NewGrid(nrows, ncols)
	return strconv.Itoa(solveDp(dp, g, nrows, ncols))
}

func Solve3(lines []string) string {
	nrows := len(lines)
	ncols := len(lines[0])
	g := createGrid2(lines, nrows, ncols)
	length := createGraph(g, nrows*5, ncols*5)
	return strconv.Itoa(length)
}

func Solve2(lines []string) string {
	nrows := len(lines)
	ncols := len(lines[0])
	g := createGrid2(lines, nrows, ncols)
	length := dijkstra(g, nrows*5, ncols*5)
	return strconv.Itoa(length)
}

func createGrid(lines []string) *Grid {
	g := make(Grid, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, col := range line {
			row = append(row, int(col-'0'))
		}
		g = append(g, row)
	}
	return &g
}

func NewGrid(rows, cols int) *Grid {
	g := make(Grid, rows)
	for i := 0; i < rows; i++ {
		g[i] = make([]int, cols)
	}
	return &g
}

func solveDp(dp *Grid, g *Grid, nrows, ncols int) int {
	(*dp)[0][0] = (*g)[0][0]
	for col := 1; col < ncols; col++ {
		(*dp)[0][col] = (*g)[0][col] + (*dp)[0][col-1]
	}
	for row := 1; row < nrows; row++ {
		(*dp)[row][0] = (*g)[row][0] + (*dp)[row-1][0]
	}
	for row := 1; row < nrows; row++ {
		for col := 0; col < ncols; col++ {
			if col > 0 {
				(*dp)[row][col] = (*g)[row][col] + utils.MinInt((*dp)[row-1][col], (*dp)[row][col-1])
			} else {
				(*dp)[row][col] = (*g)[row][col] + (*dp)[row-1][col]
			}
		}
		fmt.Println((*g)[row])
	}
	return (*dp)[nrows-1][ncols-1] - (*g)[0][0]
}


func createGrid2(lines []string, nrows, ncols int) *Grid {
	g := *createGrid(lines)
	for i := 1; i < 5; i++ {
		for nrow := 0; nrow < nrows; nrow++ {
			g = append(g, make([]int, 0))
			for ncol := 0; ncol < ncols; ncol++ {
				newVal := (g)[nrow][ncol] + i
				if newVal > 9 {
					newVal = newVal % 10 + 1
				}
				(g)[len(g)-1] = append((g)[len(g)-1], newVal)
			}
		}
	}

	for i := 0; i < len(g); i++ {
		for k := 1; k < 5; k++ {
			for j := 0; j < ncols; j++ {
				newVal := g[i][j] + k
				if newVal > 9 {
					newVal = newVal % 10 + 1
				}
				g[i] = append(g[i], newVal)
			}
		}
	}
	fmt.Println(len(g), len(g[0]))
	return &g
}

type Loc [2]int

func encode(row, col, nrows int) int {
	return row * nrows + col
}

func dijkstra(g *Grid, nrows, ncols int) int {
	visited := make([]bool, nrows*ncols)
	cost := make([]int, nrows*ncols)
	for i := 1; i < nrows*ncols; i++ {
		cost[i] = math.MaxInt
	}
	pq :=  make(utils.PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &utils.Item{0, 0, 0})
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for pq.Len() > 0 {
		head := heap.Pop(&pq).(*utils.Item)

		headRow, headCol := decode(head.Node, ncols)
		headPos := encode(headRow, headCol, nrows)
		visited[headPos] = true

		for _, direction := range directions { // change to edges (calc uprfront)
			newx := headRow + direction[0]
			newy := headCol + direction[1]
			newPos := encode(newy, newx, nrows)
			if 0 <= newx && newx < ncols && 0 <= newy && newy < nrows && !visited[newPos] {
				old_cost := cost[newPos]
				new_cost := cost[headPos] + (*g)[newy][newx]
				if new_cost < old_cost {
					cost[newPos] = new_cost
					heap.Push(&pq, &utils.Item{newPos, new_cost, 0})
				}
			}

		}
	}
	return cost[encode(nrows-1, ncols-1, nrows)]
}

func (l *Loc) encode(ncols  int) int {
	return l[0] * ncols + l[1]
}

func decode(l int, ncols int) (int, int) {
	return l / ncols, l%ncols
}

func createGraph(g *Grid, nrows, ncols  int) (int) {
	gg := graph.New(nrows * ncols)
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			for _, direction := range directions {
				newx := j + direction[0]
				newy := i + direction[1]
				if 0 <= newx && newx < ncols && 0 <= newy && newy < nrows {
					gg.AddCost((&Loc{i, j}).encode(ncols), (&Loc{newy, newx}).encode(ncols), int64((*g)[newy][newx]))
					gg.AddCost((&Loc{newy, newx}).encode(ncols), (&Loc{i, j}).encode(ncols), int64((*g)[i][j]))
				}
			}
		}
	}
	_, dist :=  graph.ShortestPath(gg, 0, (&Loc{nrows-1, ncols-1}).encode(ncols))
	return int(dist)
}
