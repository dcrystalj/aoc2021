package day24

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day24.1", Solve)
	registrator.Register("day24.2", Solve2)
	registrator.Register("day24.3", Solve3)

}

func Solve(lines []string) string {
	commands := ReadCommands(lines)
	res, _ := dfs(DfsState{State{0, 0, 0, 0}, 0}, commands, DfsMemo{}, false)
	return utils.ReverseString(fmt.Sprint(res))
}

func Solve2(lines []string) string {
	commands := ReadCommands(lines)
	res, _ := dfs(DfsState{State{0, 0, 0, 0}, 0}, commands, DfsMemo{}, true)
	return utils.ReverseString(fmt.Sprint(res))
}

// 3x slower
func Solve3(lines []string) string {
	commands := ReadCommands(lines)
	res, _ := dfs1(DfsState{State{0, 0, 0, 0}, 0}, commands, DfsMemo{}, true)
	return utils.ReverseString(fmt.Sprint(res))
}

type CommandInput struct {
	a byte
}

type CommandLiteral struct {
	name string
	a    byte
	b    int
}

type CommandVar struct {
	name string
	a    byte
	b    byte
}

type Command interface{}

func ReadCommand(line string) interface{} {
	splitted := strings.Split(line, " ")
	if splitted[0] == "inp" {
		return &CommandInput{splitted[1][0]}
	}
	value, err := strconv.Atoi(splitted[2])
	if err == nil {
		return &CommandLiteral{splitted[0], splitted[1][0], value}
	} else {
		return &CommandVar{splitted[0], splitted[1][0], splitted[2][0]}
	}
}

func ReadCommands(lines []string) []interface{} {
	list := make([]interface{}, 0)
	for _, line := range lines {
		list = append(list, ReadCommand(line))
	}
	return list
}

type State [4]int
type Option [14]int
type DfsState struct {
	s State
	c int
}
type DfsMemo map[DfsState]DfsRes
type DfsRes struct {
	res   int64
	found bool
}

func dfs(s DfsState, commands []interface{}, memo DfsMemo, reversed bool) (int64, bool) {
	memoValue, memoHas := memo[s]
	if memoHas {
		return memoValue.res, memoValue.found
	}
	nums := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if reversed {
		nums = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
Loop:
	for _, i := range nums {
		localState := s.s
		localState.read(i)
		for k := s.c + 1; k < len(commands); k++ {
			switch commands[k].(type) {
			case *CommandInput:
				localDfsState := DfsState{localState, k}
				res, found := dfs(localDfsState, commands, memo, reversed)
				if found {
					newRes := res*10 + int64(i)
					memo[localDfsState] = DfsRes{res, found}
					return newRes, found
				} else {
					continue Loop
				}
			default:
				localState.execute(commands[k])
			}
		}
		if localState.isValid() {
			memo[DfsState{localState, len(commands)}] = DfsRes{int64(i), true}
			return int64(i), true
		}
	}
	memo[s] = DfsRes{0, false}
	return 0, false
}

// slower version
func dfs1(state DfsState, commands []interface{}, memo DfsMemo, reversed bool) (int64, bool) {
	nums := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if reversed {
		nums = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	for i := state.c; i < len(commands); i++ {
		switch commands[i].(type) {
		case *CommandInput:
			for _, j := range nums {
				state.s.read(j)
				newState := DfsState{state.s, i + 1}
				memoValue, memoHas := memo[newState]
				if memoHas {
					return memoValue.res, memoValue.found
				}
				res, found := dfs1(newState, commands, memo, reversed)
				if found {
					newRes := res*10 + int64(j)
					return newRes, found
				}
				memo[newState] = DfsRes{res, found}
			}
		default:
			state.s.execute(commands[i])
		}
	}

	if state.s.isValid() {
		memo[state] = DfsRes{0, true}
		return 0, true
	} else {
		memo[state] = DfsRes{0, false}
		return 0, false
	}
}

func (s State) isValid() bool {
	return s[3] == 0
}

func (s *State) execute(command interface{}) {
	switch command.(type) {
	case *CommandVar:
		switch command.(*CommandVar).name {
		case "mul":
			s[command.(*CommandVar).a-'w'] *= s[command.(*CommandVar).b-'w']
		case "add":
			s[command.(*CommandVar).a-'w'] += s[command.(*CommandVar).b-'w']
		case "mod":
			s[command.(*CommandVar).a-'w'] %= s[command.(*CommandVar).b-'w']
		case "div":
			s[command.(*CommandVar).a-'w'] /= s[command.(*CommandVar).b-'w']
		case "eql":
			if s[command.(*CommandVar).a-'w'] == s[command.(*CommandVar).b-'w'] {
				s[command.(*CommandVar).a-'w'] = 1
			} else {
				s[command.(*CommandVar).a-'w'] = 0
			}
		}
	case *CommandLiteral:
		switch command.(*CommandLiteral).name {
		case "mul":
			s[command.(*CommandLiteral).a-'w'] *= command.(*CommandLiteral).b
		case "add":
			s[command.(*CommandLiteral).a-'w'] += command.(*CommandLiteral).b
		case "mod":
			s[command.(*CommandLiteral).a-'w'] %= command.(*CommandLiteral).b
		case "div":
			s[command.(*CommandLiteral).a-'w'] /= command.(*CommandLiteral).b
		case "eql":
			if s[command.(*CommandLiteral).a-'w'] == command.(*CommandLiteral).b {
				s[command.(*CommandLiteral).a-'w'] = 1
			} else {
				s[command.(*CommandLiteral).a-'w'] = 0
			}
		}
	default:
		panic("Invalid state")
	}
}

func (s *State) read(i int) {
	s[0] = i
}
