package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)


func init() {
    registrator.Register("day5.1", Solve)
    registrator.Register("day5.2", Solve2)
}

const BOARD_SIZE = 1000

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func makeLine(x1 int, y1 int, x2 int, y2 int) Line {
	return Line{x1: x1, y1: y1, x2: x2, y2: y2}
}

func getLines(rows []string) []Line{
	lines := make([]Line, 0)
	for _, row := range rows {
		points := strings.Split(row, " -> ")
		p1 := utils.StringToInts(points[0], ",")
		p2 := utils.StringToInts(points[1], ",")
		lines = append(lines, makeLine(p1[0], p1[1], p2[0], p2[1]))
	}
	return lines
}

func drawLine(board *[BOARD_SIZE][BOARD_SIZE]int, line Line) {
	startx := line.x1
	endx := line.x2
	starty := line.y1
	endy := line.y2
	dx := 1
	dy := 1
	if line.x1 == line.x2 {
		dx = 0
	}
	if line.x1 > line.x2 {
		dx = -1
	}
	if line.y2 == line.y1 {
		dy = 0
	}
	if line.y1 > line.y2 {
		dy = -1
	}
	for startx != endx || starty != endy {
		board[starty][startx] += 1
		if startx != endx {
			startx += dx
		}
		if starty != endy {
			starty += dy
		}
	}
	board[starty][startx] += 1
}

func calcOverlap(board *[BOARD_SIZE][BOARD_SIZE]int) int {
	cnt := 0
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if board[i][j] >= 2 {
				cnt += 1
			}
		}
	}
	return cnt
}

func printBoard (board *[BOARD_SIZE][BOARD_SIZE]int) {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func Solve(rows []string) string {
	var board [BOARD_SIZE][BOARD_SIZE]int
	lines := getLines(rows)
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			drawLine(&board, line)
		}
	}
	printBoard(&board)
	return strconv.Itoa(calcOverlap(&board))
}

func Solve2(rows []string) string {
	var board [BOARD_SIZE][BOARD_SIZE]int
	lines := getLines(rows)
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 || utils.AbsDiffInt(line.x1, line.x2) == utils.AbsDiffInt(line.y1, line.y2){
			drawLine(&board, line)
		}
	}
	// printBoard(&board)
	return strconv.Itoa(calcOverlap(&board))
}
