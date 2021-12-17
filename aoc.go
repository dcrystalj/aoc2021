package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/dcrystalj/aoc2021/day1"
	_ "github.com/dcrystalj/aoc2021/day10"
	_ "github.com/dcrystalj/aoc2021/day11"
	_ "github.com/dcrystalj/aoc2021/day12"
	_ "github.com/dcrystalj/aoc2021/day13"
	_ "github.com/dcrystalj/aoc2021/day14"
	_ "github.com/dcrystalj/aoc2021/day15"
	_ "github.com/dcrystalj/aoc2021/day16"
	_ "github.com/dcrystalj/aoc2021/day17"
	_ "github.com/dcrystalj/aoc2021/day18"
	_ "github.com/dcrystalj/aoc2021/day2"
	_ "github.com/dcrystalj/aoc2021/day3"
	_ "github.com/dcrystalj/aoc2021/day4"
	_ "github.com/dcrystalj/aoc2021/day5"
	_ "github.com/dcrystalj/aoc2021/day6"
	_ "github.com/dcrystalj/aoc2021/day7"
	_ "github.com/dcrystalj/aoc2021/day8"
	_ "github.com/dcrystalj/aoc2021/day9"
	"github.com/dcrystalj/aoc2021/registrator"
)

func main() {
	fmt.Println(os.Args)
	day := os.Args[1][:5]
	fmt.Println("./" + day + "/" + os.Args[2])
	contents, _ := ioutil.ReadFile("./" + day + "/" + os.Args[2])
	lines := strings.Split(string(contents), "\n")
	lines = lines[:len(lines)-1]

	fmt.Println(registrator.Run(os.Args[1], lines))

}
