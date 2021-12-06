package day7

import (
	"math"
	"sort"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)


func init() {
    registrator.Register("day7.1", Solve)
    registrator.Register("day7.2", Solve2)
}

func Solve(lines []string) string {
	input := utils.StringToInts(lines[0], ",")
	sort.Ints(input)
	min := math.MaxInt
	for s := range(input) {
		min = utils.MinInt(getFuel(&input, s), min)
	}
	return strconv.Itoa(min)
}

func getFuel(sorted *[]int, median int) int {
	fuel := 0
	for _, val := range *sorted {
		fuel += utils.AbsDiffInt(val, (*sorted)[median])
	}
	return fuel
}

func Solve2(lines []string) string {
	input := utils.StringToInts(lines[0], ",")
	sort.Ints(input)
	min := math.MaxInt
	for i := 0; i < input[len(input) - 1]; i++ {
		min = utils.MinInt(getFuel2(&input, i), min)
	}
	return strconv.Itoa(min)
}


func getFuel2(sorted *[]int, center int) int {
	fuel := 0
	for _, val := range *sorted {
		diff := utils.AbsDiffInt(val, center)
		if (diff > 0) {
			fuel += (diff * (diff + 1)) / 2
		}
	}
	return fuel
}
