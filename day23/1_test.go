package day23

import (
	"testing"
)

func TestMoveRemoveFromCorridor(t *testing.T) {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
		},
		NewHall(),
	}
	d.move(&CorridorMove{&d.corridors[0], &d.hallPlaces[0], 1})
	if len(d.corridors[0].items) != 1 {
		t.Error("Corridor was not removed")
	}
}

func TestGenMovesFromCorridors(t *testing.T) {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
		},
		NewHall(),
	}
	moves := d.genMovesFromCorridors()
	if len(moves) != 28 {
		t.Errorf("Should generate %d moves", len(moves))
	}
}

func TestMoveCorridor(t *testing.T) {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
		},
		NewHall(),
	}
	move := CorridorMove{&d.corridors[0], &d.hallPlaces[0], 1}
	itype := d.move(&move)

	if itype != B {
		t.Errorf("Invalid type")
	}
	if d.serialize() != "1-0,4-1,4-3,4-5,4-7,4-9,4-10,;0-[0];1-[3 2];2-[2 1];3-[0 3]" {
		t.Errorf("Should move to 123 but instead moved to %ss", d.serialize())
	}
}

func TestRevertMoveCorridor(t *testing.T) {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
		},
		NewHall(),
	}
	s1 := d.serialize()
	move := CorridorMove{&d.corridors[0], &d.hallPlaces[0], 1}
	itype := d.move(&move)
	d.revertMove(&move, itype)
	s2 := d.serialize()
	if s1 != s2 {
		t.Errorf("Should revert from %s to %s moves", s2, s1)
	}
}

func TestRevertMoveHall(t *testing.T) {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
		},
		NewHall(),
	}
	d.move(&CorridorMove{&d.corridors[0], &d.hallPlaces[0], 1})

	s1 := d.serialize()
	move := HallMove{&d.hallPlaces[0], &d.corridors[0], 1}
	itype := d.move(&move)
	d.revertMove(&move, itype)
	s2 := d.serialize()
	if s1 != s2 {
		t.Errorf("Should revert from %s to %s moves", s2, s1)
	}
}

func TestCopy(t *testing.T) {
	d := Diagram{
		[]Corridor{
			{position: 2, itype: A, maxSize: 2, items: []Itype{A, B}},
			{position: 4, itype: B, maxSize: 2, items: []Itype{D, C}},
			{position: 6, itype: C, maxSize: 2, items: []Itype{C, B}},
			{position: 8, itype: D, maxSize: 2, items: []Itype{A, D}},
		},
		NewHall(),
	}
	ds := d.serialize()
	states := []Diagram{d}
	copied := states[0].copy()
	copied.move(&CorridorMove{&copied.corridors[0], &copied.hallPlaces[0], 1})

	if ds != states[0].serialize() {
		t.Errorf("Not copied")
	}
}
