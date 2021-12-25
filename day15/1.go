package day15

import (
	"fmt"
	"strconv"

	"container/heap"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
	"github.com/yourbasic/graph"
)

func init() {
	registrator.Register("day15.1", Solve3)
	registrator.Register("day15.2", Solve4)
}

type Grid [][]int

// invalid. it can move to all directions, still passed test
func Solve(lines []string) string {
	nrows := len(lines)
	ncols := len(lines[0])
	g := createGrid(lines)
	dp := NewGrid(nrows, ncols)
	return strconv.Itoa(solveDp(dp, g, nrows, ncols))
}

// second part using graph lib
func Solve2(lines []string) string {
	nrows := len(lines)
	ncols := len(lines[0])
	g := createGrid2(lines, nrows, ncols)
	length := createGraph(g, nrows*5, ncols*5)
	return strconv.Itoa(length)
}

func Solve3(lines []string) string {
	nrows := len(lines)
	ncols := len(lines[0])
	g := createGrid(lines)
	length := dijkstra(g, nrows, ncols)
	return strconv.Itoa(length)
}

func Solve4(lines []string) string {
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
					newVal = newVal%10 + 1
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
					newVal = newVal%10 + 1
				}
				g[i] = append(g[i], newVal)
			}
		}
	}
	return &g
}

type Loc [2]int

func encode(row, col, nrows int) int {
	return row*nrows + col
}

func dijkstra(g *Grid, nrows, ncols int) int {
	visited := make([]bool, nrows*ncols)
	pq := utils.PriorityQueue{&utils.Item{Node: 0, Priority: 0, Index: 0}}
	heap.Init(&pq)
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for pq.Len() > 0 {
		head := heap.Pop(&pq).(*utils.Item)

		headRow, headCol := decode(head.Node, ncols)
		if visited[head.Node] {
			continue
		}
		if headRow == nrows-1 && headCol == ncols-1 {
			return head.Priority
		}
		visited[head.Node] = true

		for _, direction := range directions {
			newy := headRow + direction[0]
			newx := headCol + direction[1]
			newPos := encode(newy, newx, nrows)
			if 0 <= newx && newx < ncols && 0 <= newy && newy < nrows && !visited[newPos] {
				heap.Push(&pq, &utils.Item{newPos, head.Priority + (*g)[newy][newx], 0})
			}
		}
	}
	panic("Path not found")
}

func (l *Loc) encode(ncols int) int {
	return l[0]*ncols + l[1]
}

func decode(l int, ncols int) (int, int) {
	return l / ncols, l % ncols
}

func createGraph(g *Grid, nrows, ncols int) int {
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
	_, dist := graph.ShortestPath(gg, 0, (&Loc{nrows - 1, ncols - 1}).encode(ncols))
	return int(dist)
}
