package day2

import (
	"strconv"
	"strings"

	"github.com/dcrystalj/aoc2021/registrator"
)


func init() {
    registrator.Register("day2.1", Solve3)
    registrator.Register("day2.2", Solve4)
}

func Solve3(lines []string) string {
	var horizontal, depth int = 0, 0

    for _, s := range lines {
		row := strings.Split(s, " ")
		if len(row) < 2 {
			break
		}
		value, _ := strconv.Atoi(row[1])
		switch row[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
    }

    return strconv.Itoa(depth * horizontal)
}


func Solve4(lines []string) string {
	var horizontal, depth, aim int = 0, 0, 0

    for _, s := range lines {
		row := strings.Split(s, " ")
		if len(row) < 2 {
			break
		}
		value , _ := strconv.Atoi(row[1])
		switch row[0] {
		case "forward":
			horizontal += value
			depth += (value * aim)
		case "down":
			aim += value
		case "up":
			aim -= value
		}
    }

    return strconv.Itoa(depth * horizontal)
}
