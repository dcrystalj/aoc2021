package day18

import (
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
	"github.com/gammazero/deque"
)

func init() {
	registrator.Register("day18.1", Solve)
	registrator.Register("day18.2", Solve2)

}

func Solve(lines []string) string {
	stack := sumLines(lines)
	mag := magnitude(stack)
	return strconv.Itoa(mag)
}

func Solve2(lines []string) string {
	maxMagnitude := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			s1 := sumLines([]string{lines[i], lines[j]})
			s2 := sumLines([]string{lines[j], lines[i]})
			maxMagnitude = utils.MaxInt(magnitude(s1), maxMagnitude)
			maxMagnitude = utils.MaxInt(magnitude(s2), maxMagnitude)
		}
	}
	return strconv.Itoa(maxMagnitude)
}

func createStack(line string) *deque.Deque {
	d := deque.New()
	for _, c := range line {
		switch c {
		case ',':
			continue
		case '[', ']':
			d.PushBack(byte(c))
		default:
			d.PushBack(int(c - '0'))
		}
	}
	return d
}

func sumLines(lines []string) *deque.Deque {
	s1 := createStack(lines[0])
	for _, line := range lines[1:] {
		s2 := createStack(line)
		s1 = addStacks(s1, s2)
		s1 = optimize(s1)
	}
	return s1
}

func addStacks(s1, s2 *deque.Deque) *deque.Deque {
	s1.PushFront(byte('['))
	for s2.Len() > 0 {
		s1.PushBack(s2.PopFront())
	}
	s1.PushBack(byte(']'))
	return s1
}

func optimize(s *deque.Deque) *deque.Deque {
Loop:
	parenthesisCnt := 0
	// fmt.Println("loop ", serializeStack(s))
	for i := 0; i < s.Len(); i++ {
		switch s.At(i).(type) {
		case uint8:
			if s.At(i).(byte) == '[' {
				parenthesisCnt += 1
			} else if s.At(i).(byte) == ']' {
				parenthesisCnt -= 1
			}
		}

		if parenthesisCnt >= 5 {
			s = explode(s, i)
			goto Loop
		}
	}

	for i := 0; i < s.Len(); i++ {
		switch s.At(i).(type) {
		case int:
			if s.At(i).(int) > 9 {
				s = split(s, i)
				goto Loop
			}
		}
	}
	return s
}

func serializeStack(s *deque.Deque) string {
	r := ""
	for i := 0; i < s.Len(); i++ {
		left := s.At(i)
		switch left.(type) {
		case int:
			r += strconv.Itoa(left.(int)) + " "
		case byte:
			r += string(left.(byte))
		}
	}
	return r
}

func explode(s *deque.Deque, i int) *deque.Deque {
	// fmt.Println("explode", serializeStack(s), i)
	leftStack := deque.New(i)
	for j := 0; j < i; j++ {
		leftStack.PushBack(s.PopFront())
	}

	s.PopFront()
	leftNum := s.PopFront().(int)
	rightNum := s.PopFront().(int)
	s.PopFront()

	for j := leftStack.Len() - 1; j >= 0; j-- {
		switch leftStack.At(j).(type) {
		case int:
			leftStack.Set(j, leftStack.At(j).(int)+leftNum)
			j = 0
		}
	}

	for j := 0; j < s.Len(); j++ {
		switch s.At(j).(type) {
		case int:
			s.Set(j, s.At(j).(int)+rightNum)
			j = s.Len()
		}
	}

	leftStack.PushBack(0)

	for s.Len() > 0 {
		leftStack.PushBack(s.PopFront())
	}

	return leftStack
}

func split(s *deque.Deque, i int) *deque.Deque {
	leftStack := deque.New()
	for j := 0; j < i; j++ {
		leftStack.PushBack(s.PopFront())
	}
	curVal := s.PopFront().(int)
	leftStack.PushBack(byte('['))
	leftStack.PushBack(curVal / 2)
	if curVal%2 == 0 {
		leftStack.PushBack(curVal / 2)
	} else {
		leftStack.PushBack(curVal/2 + 1)
	}
	leftStack.PushBack(byte(']'))
	for s.Len() > 0 {
		leftStack.PushBack(s.PopFront())
	}
	return leftStack
}

func magnitude(s *deque.Deque) int {
Repeat:
	if s.Len() == 1 {
		return s.Front().(int)
	}
	leftStack := deque.New()
	leftStack.PushBack(s.PopFront())
Label:
	for s.Len() > 0 {
		new := s.PopFront()
		last := leftStack.Back()
		switch last.(type) {
		case int:
			switch new.(type) {
			case int:
				last = leftStack.PopBack()
				tmp := new.(int) * 2
				tmp += last.(int) * 3
				leftStack.PopBack() // remove [
				s.PopFront()        // remove ]
				leftStack.PushBack(tmp)
				goto Label
			}
		}
		leftStack.PushBack(new)
	}

	if leftStack.Len() > 0 {
		s = leftStack
		goto Repeat
	}
	return -1
}
