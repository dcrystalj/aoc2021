package day4

import (
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)


func init() {
    registrator.Register("day4.1", Solve)
    registrator.Register("day4.2", Solve2)
}

func getBingos(lines []string) [][][]int {
	b, i, j := 0,0,0
	var bingos [][][]int = make([][][]int, (len(lines)+1) / 6)
	bingos[0] = make([][]int, 5)
	for _, line := range lines {
		if line == "" {
			b += 1
			i = 0
			j = 0
			bingos[b] = make([][]int, 5)
			continue
		}
		bingos[b][i] = make([]int, 5)
		bingoRow := utils.StringToInts(line, " ")
		for _, val := range bingoRow {
			bingos[b][i][j] = val
			j += 1
		}
		j = 0
		i += 1
	}
	return bingos
}

func crossNumber(bingos [][][]int, number int) {
	for x, bingo := range bingos {
		for i, row := range bingo {
			for j, col:= range row {
				if col == number {
					bingos[x][i][j] = -1
				}
			}
		}
	}
}

func isBingoRowCol(bingo [][]int) bool {
	for i := 0; i < 5; i++ {
		isBingo := true
		for j := 0; j < 5; j++ {
			if bingo[i][j] != -1 {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		isBingo := true
		for j := 0; j < 5; j++ {
			if bingo[j][i] != -1 {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}
	return false
}
func getSolvedBingo(bingos [][][]int, toSkip map[int]bool) (int, bool) {
	for x, bingo := range bingos {
		_, exist := toSkip[x]
		if (exist) {
			continue
		}
		if isBingoRowCol(bingo) {
			return x, true
		}
	}
	return -1, false
}

func sumNonCrossed(bingo [][] int) int {
	cnt := 0
	for _, i := range bingo {
		for _, j := range i {
			if j != -1 {
				cnt += j
			}
		}
	}
	return cnt
}

func Solve(lines []string) string {
	inputs := utils.StringToInts(lines[0], ",")
	bingos := getBingos(lines[2:])
	toSkip := make(map[int]bool)
	for _, in := range inputs {
		crossNumber(bingos, in)
		index, wasSolved := getSolvedBingo(bingos, toSkip)
		if (wasSolved) {
			utils.PrintArray(bingos[index])
			return strconv.Itoa(int(sumNonCrossed(bingos[index])) * int(in))
		}
	}
	return "-1"
}

func Solve2(lines []string) string {
	inputs := utils.StringToInts(lines[0], ",")
	bingos := getBingos(lines[2:])
	lastSolved, lastNumber := -1, -1
	toSkip := make(map[int]bool)
	for _, in := range inputs {
		crossNumber(bingos, in)
		index, wasSolved := getSolvedBingo(bingos, toSkip)
		for wasSolved {
			toSkip[index] = true
			lastSolved, lastNumber = index, in
			index, wasSolved = getSolvedBingo(bingos, toSkip)
		}
		if len(toSkip) == len(bingos) {
			break
		}
	}
	return strconv.Itoa(int(sumNonCrossed(bingos[lastSolved])) * int(lastNumber))
}
