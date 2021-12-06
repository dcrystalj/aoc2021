package day6

import (
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)


func init() {
    registrator.Register("day6.1", Solve)
    registrator.Register("day6.2", Solve2)
}

func getNextState(state [9]int) [9]int {
	var tmp [9]int
	tmp[8] += state[0]
	tmp[6] += state[0]
	for i := 0; i < 8; i++ {
		tmp[i] += state[i+1]
	}
	return tmp

}

func inputsToState(inputs []int) [9]int {
	var state [9]int
	for _, in := range inputs {
		state[in] += 1
	}
	return state
}

func Solve(lines []string) string {
	inputs := utils.StringToInts(lines[0], ",")
	state := inputsToState(inputs)
	for i := 0; i < 80; i++ {
		state = getNextState(state)
	}
	return strconv.Itoa(utils.Sum(state[:]))
}

func Solve2(lines []string) string {
	inputs := utils.StringToInts(lines[0], ",")
	state := inputsToState(inputs)
	for i := 0; i < 256; i++ {
		state = getNextState(state)
	}
	return strconv.Itoa(utils.Sum(state[:]))
}
