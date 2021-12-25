package day19

import (
	"math"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
	"github.com/gammazero/deque"
)

func init() {
	registrator.Register("day19.1", Solve)
	registrator.Register("day19.2", Solve2)

}

func Solve(lines []string) string {
	lines = append(lines, "")
	scanners := linesToScanner(lines)
	beacons, _ := bfs(scanners)
	return strconv.Itoa(len(beacons))
}

func Solve2(lines []string) string {
	lines = append(lines, "")
	scanners := linesToScanner(lines)
	_, positions := bfs(scanners)

	return strconv.Itoa(furthestManhattan(positions))
}

func furthestManhattan(scan []ScannerPos) int {
	dist := math.MinInt
	for i := 0; i < len(scan); i++ {
		for j := i + 1; j < len(scan); j++ {
			dist = utils.MaxInt(scan[i].manhattan(scan[j]), dist)
		}
	}
	return dist
}

func (b1 ScannerPos) manhattan(b2 ScannerPos) int {
	return utils.AbsDiffInt(b1[0], b2[0]) + utils.AbsDiffInt(b1[1], b2[1]) + utils.AbsDiffInt(b1[2], b2[2])
}

type Beacon [3]int

type Scan [3]int
type ScannerPos [3]int

type Scanner struct {
	id    int
	scans []Scan
}

func bfs(scanners []Scanner) ([]Beacon, []ScannerPos) {
	allBeaconsSet := map[Beacon]bool{}
	visited := map[int]bool{}
	q := deque.New()
	positions := []ScannerPos{{0, 0, 0}}
	q.PushBack(scanners[0].id)

	for _, scan := range scanners[0].scans {
		allBeaconsSet[Beacon{scan[0], scan[1], scan[2]}] = true
	}

	for q.Len() > 0 {
		head := q.PopBack().(int)
		if visited[head] {
			continue
		}
		visited[head] = true

		for i := 0; i < len(scanners); i++ {
			if visited[i] {
				continue
			}
			isOverlap, orientation := scanners[head].overlaps(&scanners[i])
			if isOverlap {
				tmpScan := Scanner{}
				tmpScan.id = i
				for _, scan := range scanners[i].scans {
					s := Scan{
						scan[orientation.axis[0]]*orientation.sign[0] + orientation.distance[0],
						scan[orientation.axis[1]]*orientation.sign[1] + orientation.distance[1],
						scan[orientation.axis[2]]*orientation.sign[2] + orientation.distance[2],
					}
					tmpScan.scans = append(tmpScan.scans, s)
					scanners[i] = tmpScan
					allBeaconsSet[Beacon{s[0], s[1], s[2]}] = true
					positions = append(positions, ScannerPos{
						orientation.distance[0],
						orientation.distance[1],
						orientation.distance[2],
					})
				}
				q.PushBack(i)
			}
		}
	}

	allBeacons := []Beacon{}
	for beacon := range allBeaconsSet {
		allBeacons = append(allBeacons, beacon)
	}
	return allBeacons, positions
}

func CreateScanner(id int, in []string) *Scanner {
	scans := make([]Scan, 0)
	for _, line := range in {
		parsedLine := utils.StringToInts(line, ",")
		scans = append(scans, Scan{parsedLine[0], parsedLine[1], parsedLine[2]})
	}

	return &Scanner{id, scans}
}

func linesToScanner(lines []string) []Scanner {
	prev := 0
	id := 0
	s := make([]Scanner, 0)
	for i := 1; i < len(lines); i++ {
		prev = i
		for j := i; j < len(lines); j++ {
			if lines[j] == "" {
				s = append(s, *CreateScanner(id, lines[prev:j]))
				id += 1
				i = j + 1
				break
			}
		}
	}
	return s
}

type Orientation struct {
	axis     [3]int
	sign     [3]int
	distance [3]int
}

func (s1 *Scanner) overlaps(s2 *Scanner) (bool, Orientation) {
	o := Orientation{}
	overlapPosition := [3]bool{}
AxisLoop:
	for s1axis := 0; s1axis < 3; s1axis++ {
		for _, sign := range [2]int{-1, 1} {
			for s2axis := 0; s2axis < 3; s2axis++ {

				distanceCnt := map[int]int{}
				for _, sc1 := range s1.scans {
					for _, sc2 := range s2.scans {
						distanceCnt[sc1[s1axis]-sign*sc2[s2axis]] += 1
					}
				}

				for distance, cnt := range distanceCnt {
					if cnt >= 12 {
						overlapPosition[s1axis] = true
						o.axis[s1axis] = s2axis
						o.sign[s1axis] = sign
						o.distance[s1axis] = distance
						continue AxisLoop
					}
				}
			}
		}
	}
	doesOverlap := overlapPosition[0] && overlapPosition[1] && overlapPosition[2]
	return doesOverlap, o
}

func (s1 *Scan) add(s2 *Scan) *Scan {
	return &Scan{s1[0] + s2[0], s1[1] + s2[1], s1[2] + s2[2]}
}

func (s *Scan) Equals(s2 *Scan) bool {
	for i := 0; i < 3; i++ {
		if s[i] != s2[i] {
			return false
		}
	}
	return true
}

// func (s *Scan) orientation(i int) *Scan {
// 	scans := []Scan{
// 		{s.x, s.y, s.z},
// 		{s.x, -s.y, -s.z},
// 		{-s.x, s.y, -s.z},
// 		{-s.x, -s.y, s.z},

// 		{s.x, s.z, -s.y},
// 		{s.x, -s.z, s.y},
// 		{-s.x, s.z, s.y},
// 		{-s.x, -s.z, -s.y},

// 		{s.y, s.z, s.x},
// 		{s.y, -s.z, -s.x},
// 		{-s.y, s.z, -s.x},
// 		{-s.y, -s.z, s.x},

// 		{s.y, s.x, -s.z},
// 		{s.y, -s.x, s.z},
// 		{-s.y, s.x, s.z},
// 		{-s.y, -s.x, -s.z},

// 		{s.z, s.x, s.y},
// 		{s.z, -s.x, -s.y},
// 		{-s.z, s.x, -s.y},
// 		{-s.z, -s.x, s.y},

// 		{s.z, s.y, -s.x},
// 		{s.z, -s.y, s.x},
// 		{-s.z, s.y, s.x},
// 		{-s.z, -s.y, -s.x},
// 	}
// 	return &scans[i]
// }
