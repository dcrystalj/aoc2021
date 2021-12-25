package day23

import (
	"container/heap"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day23.1", Solve)
	registrator.Register("day23.2", Solve2)

}

func Solve(lines []string) string {
	// sample
	// d := Diagram{
	// 	[]Corridor{
	// 		{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
	// 		{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
	// 		{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
	// 		{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
	// 	},
	// 	NewHall(),
	// }
	// solution := Diagram{
	// 	[]Corridor{
	// 		{position: 2, itype: A, maxSize: 2, items: []Itype{A, A}},
	// 		{position: 4, itype: B, maxSize: 2, items: []Itype{B, B}},
	// 		{position: 6, itype: C, maxSize: 2, items: []Itype{C, C}},
	// 		{position: 8, itype: D, maxSize: 2, items: []Itype{D, D}},
	// 	},
	// 	NewHall(),
	// }
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{B, C}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{A, D}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{B, D}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{C, A}},
		},
		NewHall(),
	}
	solution := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, A}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{B, B}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, C}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{D, D}},
		},
		NewHall(),
	}
	res := d.minEnergyToSolve(solution.serialize())
	return strconv.Itoa(res)
}

func Solve2(lines []string) string {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 4, items: []Itype{B, D, D, C}},
			{position: 4, itype: B, maxSize: 4, items: []Itype{A, B, C, D}},
			{position: 6, itype: C, maxSize: 4, items: []Itype{B, A, B, D}},
			{position: 8, itype: D, maxSize: 4, items: []Itype{C, C, A, A}},
		},
		NewHall(),
	}
	solution := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 4, items: []Itype{A, A, A, A}},
			{position: 4, itype: B, maxSize: 4, items: []Itype{B, B, B, B}},
			{position: 6, itype: C, maxSize: 4, items: []Itype{C, C, C, C}},
			{position: 8, itype: D, maxSize: 4, items: []Itype{D, D, D, D}},
		},
		NewHall(),
	}
	res := d.minEnergyToSolve(solution.serialize())
	return strconv.Itoa(res)
}

type Itype int

const (
	A Itype = iota
	B
	C
	D
	EMPTY
)

var energy = map[Itype]int{
	A: 1, B: 10, C: 100, D: 1000,
}

type Corridor struct {
	position int
	itype    Itype
	maxSize  int
	items    []Itype
}

type HallPlace struct {
	itype    Itype
	position int
}

func NewHall() []HallPlace {
	return []HallPlace{
		{EMPTY, 0},
		{EMPTY, 1},
		{EMPTY, 3},
		{EMPTY, 5},
		{EMPTY, 7},
		{EMPTY, 9},
		{EMPTY, 10},
	}
}

type Diagram struct {
	corridors  []Corridor
	hallPlaces []HallPlace
}

func (d1 *Diagram) copy() *Diagram {
	d2 := &Diagram{}
	for _, corridor := range d1.corridors {
		items := corridor.items
		d2.corridors = append(d2.corridors, Corridor{corridor.position, corridor.itype, corridor.maxSize, items})
	}
	// d2.hallPlaces = append(d2.hallPlaces, d1.hallPlaces...)
	for _, hp := range d1.hallPlaces {
		d2.hallPlaces = append(d2.hallPlaces, HallPlace{hp.itype, hp.position})
	}
	return d2
}

func (d *Diagram) serialize() (r string) {
	// for _, hall := range d.hallPlaces {
	// 	r += fmt.Sprintf("%d-%d,", hall.itype, hall.position)
	// }
	// for _, c := range d.corridors {
	// 	r += fmt.Sprintf(";%d-%v", c.itype, c.items)
	// }
	// return r
	// visualised
	arr := []byte{}
	for i := 0; i < 11; i++ {
		arr = append(arr, '.')
	}
	for _, hp := range d.hallPlaces {
		if hp.itype == EMPTY {
			arr[hp.position] = '.'
		} else {
			arr[hp.position] = byte('A' + int(hp.itype))
		}
	}
	for i := 0; i < d.corridors[0].maxSize; i++ {
		arr = append(arr, '#', '#', '\n', '#', '#')
		for _, c := range d.corridors {
			if len(c.items) >= c.maxSize-i {
				arr = append(arr, byte('A'+c.items[c.maxSize-i-1]), '#')
			} else {
				arr = append(arr, '.', '#')
			}
		}
	}
	arr = append(arr, '#', '#', '\n')
	r = string(arr)
	return
}

func (d *Diagram) isSolved() bool {
	for _, corridor := range d.corridors {
		if !corridor.isSolved() {
			return false
		}
	}
	for _, place := range d.hallPlaces {
		if place.itype != EMPTY {
			return false
		}
	}
	return true
}

func (c *Corridor) isSolved() bool {
	for _, item := range c.items {
		if item != c.itype {
			return false
		}
	}
	return true
}

func (d *Diagram) minEnergyToSolve(finalState string) int {
	states := []Diagram{*d}
	serializedStates := []string{states[0].serialize()}
	// visited := map[string]bool{}
	visited := map[string]bool{}
	pq := utils.PriorityQueue{&utils.Item{0, 0, 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		head := heap.Pop(&pq).(*utils.Item)

		_, has := visited[serializedStates[head.Node]]
		if has {
			continue
		}
		if serializedStates[head.Node] == finalState {
			return head.Priority
		}
		visited[serializedStates[head.Node]] = true
		state := states[head.Node].copy()
		// fmt.Println(serializedStates[head.Node])
		moves := state.genMovesFromCorridors()
		for m, move := range moves {
			t := state.move(moves[m])
			serializedState := (state).serialize()
			// fmt.Println(serializedState)
			_, has := visited[serializedState]
			if !has {
				states = append(states, *state.copy())
				serializedStates = append(serializedStates, serializedState)
				heap.Push(&pq, &utils.Item{len(states) - 1, head.Priority + move.cost, 0})
			}
			state.revertMove(moves[m], t)
		}

		moves2 := state.genMovesFromHall()
		// fmt.Println(serializedStates[head.Node])
		for m, move := range moves2 {
			t := state.move(moves2[m])
			serializedState := (state).serialize()
			// fmt.Println(serializedState)
			_, has := visited[serializedState]
			if !has {
				states = append(states, *state.copy())
				serializedStates = append(serializedStates, serializedState)
				heap.Push(&pq, &utils.Item{len(states) - 1, head.Priority + move.cost, 0})
			}
			state.revertMove(moves2[m], t)
		}
	}
	panic("Path not found")
}

func (d *Diagram) move(m interface{}) (t Itype) {
	switch (m).(type) {
	case *CorridorMove:
		cm := (m).(*CorridorMove)
		items := &cm.from.items
		t = (*items)[len(*items)-1]
		cm.to.itype = t
		*items = (*items)[:len(*items)-1]
	case *HallMove:
		hm := (m).(*HallMove)
		t = hm.from.itype
		items := &hm.to.items
		*items = append(*items, hm.from.itype)
		hm.from.itype = EMPTY
	}
	return
}

func (d *Diagram) revertMove(m interface{}, t Itype) {
	switch m.(type) {
	case *CorridorMove:
		cm := m.(*CorridorMove)
		cm.to.itype = t
		d.move(&HallMove{cm.to, cm.from, cm.cost})
	case *HallMove:
		hm := m.(*HallMove)
		d.move(&CorridorMove{hm.to, hm.from, hm.cost})
	}
}

type CorridorMove struct {
	from *Corridor
	to   *HallPlace
	cost int
}

func (d *Diagram) genMovesFromCorridors() []*CorridorMove {
	r := (make([]*CorridorMove, 0))
	for c, corridor := range d.corridors {
		corridorLen := len(corridor.items)
		// empty corridor
		if corridorLen == 0 {
			continue
		}

		// cost from corridor
		moveLength := (corridor.maxSize - corridorLen)

		// find starting hall place
		startHall := 1
		for corridor.position > d.hallPlaces[startHall].position {
			startHall++
		}

		// search left until empty space
		for i := startHall - 1; i >= 0; i-- {
			if d.hallPlaces[i].itype == EMPTY {
				e, found := energy[corridor.items[corridorLen-1]]
				if !found {
					panic("invalid energy")
				}
				cost := (moveLength + 1 + (corridor.position - d.hallPlaces[i].position)) * e
				if cost < 0 {
					panic("Invalid cost")
				}
				r = append(r, &CorridorMove{&d.corridors[c], &d.hallPlaces[i], cost})
			} else {
				break
			}
		}

		// search right until empty space
		for i := startHall; i < len(d.hallPlaces); i++ {
			if d.hallPlaces[i].itype == EMPTY {
				e, found := energy[corridor.items[corridorLen-1]]
				if !found {
					panic("invalid energy")
				}
				cost := (moveLength + 1 + (d.hallPlaces[i].position - corridor.position)) * e
				if cost < 0 {
					panic("Invalid cost")
				}
				r = append(r, &CorridorMove{&d.corridors[c], &d.hallPlaces[i], cost})
			} else {
				break
			}
		}
	}
	return r
}

type HallMove struct {
	from *HallPlace
	to   *Corridor
	cost int
}

func (d *Diagram) genMovesFromHall() []*HallMove {
	r := make([]*HallMove, 0)
	for i, hall := range d.hallPlaces {
		if hall.itype == EMPTY {
			continue
		}

		for j := i - 1; j >= 0; j-- {
			for c, corridor := range d.corridors {
				if corridor.canPush(hall.itype) && d.hallPlaces[j].position < corridor.position && corridor.position < hall.position {
					e, found := energy[hall.itype]
					if !found {
						panic("invalid energy")
					}
					cost := ((hall.position - corridor.position) + 1 + (corridor.maxSize - len(corridor.items))) * e
					if cost < 0 {
						panic("Invalid cost")
					}
					r = append(r, &HallMove{&d.hallPlaces[i], &d.corridors[c], cost})
				}
			}
			if d.hallPlaces[j].itype != EMPTY {
				break
			}
		}

		for j := i + 1; j < len(d.hallPlaces); j++ {
			for c, corridor := range d.corridors {
				if corridor.canPush(hall.itype) && d.hallPlaces[j].position > corridor.position && corridor.position > hall.position {
					e, found := energy[hall.itype]
					if !found {
						panic("invalid energy")
					}
					cost := ((corridor.position - hall.position) + 1 + (corridor.maxSize - len(corridor.items))) * e
					if cost < 0 {
						panic("Invalid cost")
					}
					r = append(r, &HallMove{&d.hallPlaces[i], &d.corridors[c], cost})
				}
			}
			if d.hallPlaces[j].itype != EMPTY {
				break
			}
		}
	}

	return removeDuplicates(r)
}

type II struct {
	a, b int
}

func removeDuplicates(moves []*HallMove) []*HallMove {
	r := []*HallMove{}
	visited := map[II]bool{}
	for i, m := range moves {
		key := II{m.from.position, m.to.position}
		_, has := visited[key]
		if has {
			continue
		} else {
			visited[key] = true
			r = append(r, moves[i])
		}
	}
	return r
}

func (c *Corridor) canPush(hallType Itype) bool {
	if c.itype != hallType || len(c.items) == c.maxSize {
		return false
	}

	for _, i := range c.items {
		if i != c.itype {
			return false
		}
	}

	return true
}
