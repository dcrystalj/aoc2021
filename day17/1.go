package day17

import (
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day17.1", Solve)
	registrator.Register("day17.2", Solve)

}

type Target struct {
	x1, y1, x2, y2 int
}

func Solve(lines []string) string {
	values := utils.StringToInts(lines[0], ",")
	t := Target{values[0], values[2], values[1], values[3]}
	res1, res2 := getHeightCount(t)
	return strconv.Itoa(res1) + " " + strconv.Itoa(res2)
}

func getHeightCount(t Target) (int, int) {
	maxHeight := 0
	count := 0
	for xV := 0; xV <= t.x2; xV++ {
		for yV := -500; yV < 500; yV++ {
			xpos, ypos, vx, vy, height := 0, 0, xV, yV, 0
			if t.x1 <= xpos && xpos <= t.x2 && t.y1 <= ypos && ypos <= t.y2 {
				maxHeight = utils.MaxInt(height, maxHeight)
				count += 1
				continue
			}
			for xpos <= t.x2 && ypos >= t.y1 {
				xpos += vx
				ypos += vy
				vy -= 1
				if vx > 0 {
					vx -= 1
				} else if vx < 0 {
					vx += 1
				}

				height = utils.MaxInt(height, ypos)

				if t.x1 <= xpos && xpos <= t.x2 && t.y1 <= ypos && ypos <= t.y2 {
					maxHeight = utils.MaxInt(height, maxHeight)
					count += 1
					break
				}
			}
		}
	}
	return maxHeight, count
}
