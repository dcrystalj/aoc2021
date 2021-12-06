package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day13.1", Solve)
	registrator.Register("day13.2", Solve2)
}

type Paper map[Loc]bool
type Fold struct {
	value   int
	isXaxis bool
}

type Loc struct {
	y, x int
}

func (p *Paper) fold(f Fold) {
	for loc := range *p {
		if f.isXaxis && loc.x > f.value {
			diff := (loc.x - f.value)
			newX := f.value - diff
			(*p)[Loc{loc.y, newX}] = true
			delete(*p, loc)
		} else if !f.isXaxis && loc.y > f.value {
			diff := (loc.y - f.value)
			newY := f.value - diff
			(*p)[Loc{newY, loc.x}] = true
			delete(*p, loc)
		}
	}
}

func (p *Paper) print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 50; j++ {
			_, has := (*p)[Loc{i, j}]
			if has {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Solve(lines []string) string {
	paper, folds := createPaper(lines)
	paper.fold((*folds)[0])
	return strconv.Itoa(len(*paper))
}

func createPaper(lines []string) (*Paper, *[]Fold) {
	p := make(Paper, 0)
	foldStart := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			foldStart = i + 1
			break
		}
		splittedLine := utils.StringToInts(lines[i], ",")
		l := Loc{splittedLine[1], splittedLine[0]}
		p[l] = true
	}
	folds := make([]Fold, 0)
	for i := foldStart; i < len(lines); i++ {
		splittedLine := strings.Split(lines[i], "=")
		value, _ := strconv.Atoi(splittedLine[1])
		if splittedLine[0][len(splittedLine[0])-1:] == "x" {
			folds = append(folds, Fold{value, true})
		} else {
			folds = append(folds, Fold{value, false})
		}
	}
	return &p, &folds
}

func Solve2(lines []string) string {
	paper, folds := createPaper(lines)
	for _, fold := range *folds {
		paper.fold(fold)
	}
	(*paper).print()
	return strconv.Itoa(len(*paper))
}
