package day21

import (
	"fmt"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day21.1", Solve)
	registrator.Register("day21.2", Solve2)

}

func Solve(lines []string) string {
	d := Dice{1}
	p1 := NewPlayer(lines[0][28])
	p2 := NewPlayer(lines[1][28])
	cnt := 0
	for !(p1.hasWon() || p2.hasWon()) {
		rolled := d.roll()
		if cnt%2 == 0 {
			p1.move(rolled)
		} else {
			p2.move(rolled)
		}
		cnt++
	}
	cnt *= 3
	loser := p2.score
	if p2.hasWon() {
		loser = p1.score
	}
	return strconv.Itoa(loser * cnt)
}

func Solve2(lines []string) string {
	p1 := NewPlayer(lines[0][28])
	p2 := NewPlayer(lines[1][28])
	return fmt.Sprintf("%d", part2(*p1, *p2))
}

type Player struct {
	score, pos int
}

type Dice struct {
	state int
}

func (d *Dice) roll() (sum int) {
	sum = 0
	if d.state >= 999 {
		for i := 0; i < 3; i++ {
			if d.state+i > 1000 {
				sum += (d.state+i)%1000 + 1
			} else {
				sum += (d.state + i) % 1000
			}
		}
	} else {
		sum = (d.state + 1) * 3
	}
	d.state += 3
	if d.state > 1000 {
		d.state = d.state%1000 + 1
	}
	return
}

func NewPlayer(start byte) *Player {
	val := int(start - '0')
	return &Player{0, val}
}

func (p *Player) move(value int) {
	toMove := value % 10
	p.pos = (p.pos + toMove) % 10
	if p.pos == 0 {
		p.pos = 10
	}
	p.score += p.pos
}

func (p *Player) hasWon() bool {
	return p.score >= 1000
}

func createFrequencies() map[int]int {
	f := map[int]int{}
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				f[r1+r2+r3] += 1
			}
		}
	}
	return f
}

// type State struct {
// 	player1 Player
// 	player2 Player
// }
// func Part2(p1, p2 Player) int64 {
// 	states := map[State]int64{{p1, p2}: 1}

// 	frequencies := createFrequencies()
// 	var p1wins, p2wins int64 = 0, 0

// 	for step := 0; step <= 400000; step++ {
// 		nextStates := map[State]int64{}
// 		for state, count := range states {
// 			for roll, fCount := range frequencies {
// 				n := count * int64(fCount)

// 				if step%2 == 0 {
// 					state.player1.move(roll)
// 					if state.player1.score >= 21 {
// 						p1wins += n
// 					} else {
// 						nextStates[state] += n
// 					}
// 				} else {
// 					state.player2.move(roll)
// 					if state.player2.score >= 21 {
// 						p2wins += n
// 					} else {
// 						nextStates[state] += n
// 					}
// 				}
// 			}
// 		}
// 		states = nextStates
// 		if len(states) == 0 {
// 			return utils.MaxInt64(p1wins, p2wins)
// 		}
// 	}
// 	return -1
// }

type State struct {
	player1       Player
	player2       Player
	isPlayer1turn bool
}

func part2(p1, p2 Player) int64 {
	states := map[State]int64{{p1, p2, true}: 1}

	frequencies := createFrequencies()
	var p1wins, p2wins int64 = 0, 0

	for len(states) > 0 {
		nextStates := map[State]int64{}
		for state, stateCount := range states {
			for roll, fCount := range frequencies {
				n := stateCount * int64(fCount)
				copyState := state

				if state.isPlayer1turn {
					copyState.player1.move(roll)
					if copyState.player1.score >= 21 {
						p1wins += n
					} else {
						nextStates[State{copyState.player1, copyState.player2, false}] += n
					}
				} else {
					copyState.player2.move(roll)
					if copyState.player2.score >= 21 {
						p2wins += n
					} else {
						nextStates[State{copyState.player1, copyState.player2, true}] += n
					}
				}
			}
		}
		states = nextStates
	}
	return utils.MaxInt64(p1wins, p2wins)
}
