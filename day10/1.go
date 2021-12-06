package day10

import (
	"sort"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

var openingToClosing map[byte]byte

func init() {
    registrator.Register("day10.1", Solve)
    registrator.Register("day10.2", Solve2)

	openingToClosing = map[byte]byte{'{': '}', '(': ')', '[': ']', '<' : '>'}

}

func Solve(lines []string) string {
	sum := 0
	for _, line := range lines {
		char := getIllegalChar(line)
		sum += charToVal(char)
	}
	return strconv.Itoa(sum)
}

func getIllegalChar(s string) byte {
	stack := deque.New()
	openingToClosing := map[byte]byte{'{': '}', '(': ')', '[': ']', '<' : '>'}
	for _, char := range s {
		bchar := byte(char)
		_, found := openingToClosing[bchar]
		if found {
			stack.PushRight(bchar)
			continue
		} else if stack.Empty() {
			return bchar
		}
		expecting := openingToClosing[stack.Right().(byte)]
		if expecting != bchar {
			return bchar
		}
		stack.PopRight()
	}
	return '0'
}

func charToVal(char byte) int {
	switch char {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func Solve2(lines []string) string {
	allPoints := make([]int, 0)
	for _, line := range lines {
		stack := getUnfinishedStack(line)
		if stack.Empty() {
			continue
		}
		value := stackValue(stack)
		allPoints = append(allPoints, value)
	}
	sort.Ints(allPoints)
	return strconv.Itoa(allPoints[len(allPoints) / 2])
}

func stackValue(stack *deque.Deque) int {
	sum := 0
	for !stack.Empty() {
		char := openingToClosing[stack.Right().(byte)]
		sum *= 5
		sum += validCharToVal(char)
		stack.PopRight()
	}
	return sum
}

func validCharToVal(char byte) int {
	switch char {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}

func getUnfinishedStack(s string) *deque.Deque {
	stack := deque.New()
	for _, char := range s {
		bchar := byte(char)
		_, found := openingToClosing[bchar]
		if found {
			stack.PushRight(bchar)
			continue
		} else if stack.Empty() {
			return stack
		}
		expecting := openingToClosing[stack.Right().(byte)]
		if expecting != bchar {
			return deque.New()
		}
		stack.PopRight()
	}
	return stack
}
