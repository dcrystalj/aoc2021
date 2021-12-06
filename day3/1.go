package day3

import (
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
)


func init() {
    registrator.Register("day3.1", Solve)
    registrator.Register("day3.2", Solve2)
}

func Solve(lines []string) string {
	counter := make([]int, len(lines[0]))

    for _, s := range lines {
		for pos, char := range s {
			if char == '1' {
				counter[pos] += 1
			}
		}
    }

	n := len(lines)
	gammaRateString := ""
	for _, c := range counter {
		if (c > n/2) {
			gammaRateString += "1"
		} else {
			gammaRateString += "0"
		}
	}

	res, _ := strconv.ParseInt(gammaRateString, 2, 32)
	var shift int64 = 1 << len(lines[0])- 1
	var q int64 = -((^res-1) ^ shift)
	return strconv.Itoa(int(res * q))
}

func getTank(set map[string]bool, flagDirection int) int64 {
	position := 0
	for len(set) > 0 {
		boundary := (len(set)+1) / 2
		counter := 0
		for k := range set {
			if k[position] == '1' {
				counter +=1
			}
		}
		mostCommon := flagDirection
		if counter < boundary {
			mostCommon = 1-flagDirection
		}
		toKeep := strconv.Itoa(mostCommon)[0]

		for k := range set {
			if len(set) == 1 {
				res, _ := strconv.ParseInt(k, 2, 32)
				return res
			}
			if k[position] != byte(toKeep) {
				delete(set, k)
			}
		}
		position += 1
	}
	return -1
}

func Solve2(lines []string) string {
	set := map[string]bool{}
	set2 := map[string]bool{}
	counter := make([]int, len(lines[0]))
    for _, s := range lines {
		for pos, char := range s {
			if char == '1' {
				counter[pos] += 1
			}
		}
		if len(s) > 0 {
			set[s] = true
			set2[s] = true
		}
    }
	oxygen, co2 := getTank(set, 1), getTank(set2, 0)
	return strconv.Itoa(int(oxygen * co2))
}
