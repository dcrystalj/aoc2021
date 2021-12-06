package day12

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/dcrystalj/aoc2021/registrator"
)

func init() {
	registrator.Register("day12.1", Solve)
	registrator.Register("day12.2", Solve2)
}

func createPaths(lines []string) map[string][]string {
	paths := make(map[string][]string, 0)
	for _, line := range lines {
		splitted := strings.Split(line, "-")

		_, found := paths[splitted[0]]
		if found {
			paths[splitted[0]] = append(paths[splitted[0]], splitted[1])
		} else {
			paths[splitted[0]] = []string{splitted[1]}
		}

		_, found = paths[splitted[1]]
		if found {
			paths[splitted[1]] = append(paths[splitted[1]], splitted[0])
		} else {
			paths[splitted[1]] = []string{splitted[0]}
		}
	}
	return paths
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func dfs(pathMapper *map[string][]string, visited *map[string]bool, parent string) int {
	cnt := 0
	for _, el := range (*pathMapper)[parent] {
		if el == parent || el == "start" {
			continue
		}
		if el == "end" {
			cnt += 1
			continue
		} else if IsUpper(el) {
			cnt += dfs(pathMapper, visited, el)
		} else {
			wasVisited, found := (*visited)[el]
			if wasVisited && found {
				continue
			}
			(*visited)[el] = true
			cnt += dfs(pathMapper, visited, el)
			(*visited)[el] = false
		}
	}
	return cnt
}

func Solve(lines []string) string {
	pathMapper := createPaths(lines)
	visited := make(map[string]bool, 0)
	numWays := dfs(&pathMapper, &visited, "start")
	return strconv.Itoa(numWays)
}

func Solve2(lines []string) string {
	pathMapper := createPaths(lines)
	visited := make(map[string]int, 0)
	path = make([]string, 0)
	numWays := dfs2(&pathMapper, &visited, "start", true)
	return strconv.Itoa(numWays)
}

var path []string

func dfs2(pathMapper *map[string][]string, visited *map[string]int, parent string, canVisitTwice bool) int {
	cnt := 0
	for _, el := range (*pathMapper)[parent] {
		if el == parent || el == "start" {
			continue
		}
		path = append(path, el)
		if el == "end" {
			cnt += 1
			// fmt.Println(path)
			path = path[:len(path)-1]
			continue
		} else if IsUpper(el) {
			cnt += dfs2(pathMapper, visited, el, canVisitTwice)
		} else {
			wasVisited, found := (*visited)[el]

			if found && ((wasVisited == 1 && !canVisitTwice) || wasVisited > 1) {
				path = path[:len(path)-1]
				continue
			}
			(*visited)[el] += 1
			if wasVisited >= 1 || !canVisitTwice {
				cnt += dfs2(pathMapper, visited, el, false)
			} else {
				cnt += dfs2(pathMapper, visited, el, true)
			}
			(*visited)[el] -= 1
		}
		path = path[:len(path)-1]
	}
	return cnt
}
