package day1

import (
	"strconv"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"

	"github.com/dcrystalj/aoc2021/registrator"
)

func init() {
    registrator.Register("day1.1", Solve1)
    registrator.Register("day1.2", Solve2)
}

func Solve1(lines []string) string {
    var prevVal, counter int = 0, 0
    isFirst := true
    for _, line := range lines {
        newVal, _ := strconv.Atoi(line)
        if !isFirst && prevVal < newVal {
            counter += 1
        }
        prevVal = newVal
        isFirst = false
    }
    return strconv.Itoa(counter)
}


func Solve2(lines []string) string {
    d := deque.New()
    var prevSum, counter int = 0, 0
    for _, line := range lines {
        newVal, _ := strconv.Atoi(line)
        if d.Size() < 3 {
            d.PushRight(newVal)
            prevSum += newVal
            continue
        }
        newSum := prevSum - d.Left().(int)  + newVal
        if  prevSum < newSum {
            counter += 1
        }
        d.PopLeft()
        d.PushRight(newVal)
        prevSum = newSum
    }
    return strconv.Itoa(counter)
}
