package day22

import (
	"fmt"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day22.1", Solve)
	registrator.Register("day22.2", Solve2)

}

func Solve(lines []string) string {
	cubes := linesToCubes(lines)
	cubes = filterCubesIn50Region(cubes)
	cubes = mergeCubes(cubes)
	return fmt.Sprintf("%d", cubesOn(cubes))
}

func Solve2(lines []string) string {
	cubes := linesToCubes(lines)
	cubes = mergeCubes(cubes)
	return fmt.Sprintf("%d", cubesOn(cubes))
}

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
	isOn                   bool
}

func linesToCubes(lines []string) []Cube {
	cubes := make([]Cube, 0)
	for _, line := range lines {
		cubes = append(cubes, *parseLine(line))
	}
	return cubes
}

func parseLine(line string) *Cube {
	c := Cube{}
	desiredState := ""
	_, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &desiredState, &c.x1, &c.x2, &c.y1, &c.y2, &c.z1, &c.z2)
	if err != nil {
		fmt.Println(err)
	}
	if desiredState == "on" {
		c.isOn = true
	}
	return &c
}

func mergeCubes(cubes []Cube) []Cube {
	merged := make([]Cube, 0)
	for _, cube := range cubes {
		merged_len := len(merged)
		for i := 0; i < merged_len; i++ {
			if cube.intersects(&merged[i]) {
				merged = append(merged, *cube.createIntersection(&merged[i]))
			}
		}
		if cube.isOn {
			merged = append(merged, cube)
		}
	}
	return merged
}

func (c1 *Cube) intersects(c2 *Cube) bool {
	return !(c1.x2 < c2.x1 || c2.x2 < c1.x1 || c1.y2 < c2.y1 || c2.y2 < c1.y1 || c1.z2 < c2.z1 || c2.z2 < c1.z1)
}

func (c *Cube) size() int64 {
	sign := int64(-1)
	if c.isOn {
		sign = 1
	}
	return sign * (utils.AbsDiffInt64(c.x1, c.x2) + 1) * (utils.AbsDiffInt64(c.y1, c.y2) + 1) * (utils.AbsDiffInt64(c.z1, c.z2) + 1)
}

func (c1 *Cube) createIntersection(c2 *Cube) *Cube {
	c := Cube{
		utils.MaxInt(c1.x1, c2.x1),
		utils.MinInt(c1.x2, c2.x2),
		utils.MaxInt(c1.y1, c2.y1),
		utils.MinInt(c1.y2, c2.y2),
		utils.MaxInt(c1.z1, c2.z1),
		utils.MinInt(c1.z2, c2.z2),
		!c2.isOn}
	return &c
}

func cubesOn(cubes []Cube) int64 {
	cnt := int64(0)
	for _, c := range cubes {
		cnt += c.size()
	}
	return cnt
}

func filterCubesIn50Region(cubes []Cube) (filtered []Cube) {
	for _, c := range cubes {
		if (-50 <= c.x1 && c.x1 <= 50 && -50 <= c.x2 && c.x2 <= 50) && (-50 <= c.y1 && c.y1 <= 50 && -50 <= c.y2 && c.y2 <= 50) && (-50 <= c.z1 && c.z1 <= 50 && -50 <= c.z2 && c.z2 <= 50) {
			filtered = append(filtered, c)
		}
	}
	return
}
