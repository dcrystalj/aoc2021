package day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

func init() {
	registrator.Register("day14.1", Solve)
	registrator.Register("day14.2", Solve2)
}

func Solve(lines []string) string {
	template := lines[0]
	counter := templateToCounter(template)
	rules := NewRules(lines[2:])
	for i := 0; i < 10; i++ {
		counter = counter.makeStep(rules)
	}
	fmt.Println(counter)
	counter = counter.count()
	min, max := counter.minMax()
	return strconv.Itoa(max - min)
}

func Solve2(lines []string) string {
	template := lines[0]
	counter := templateToCounter(template)
	rules := NewRules(lines[2:])
	for i := 0; i < 40; i++ {
		counter = counter.makeStep(rules)
	}
	fmt.Println(counter)
	counter = counter.count()
	min, max := counter.minMax()
	return strconv.Itoa(max - min)
}

type Rules map[string]string
type Counter map[string]int

func NewRules(lines []string) *Rules {
	rules := make(Rules)
	for _, line := range lines {
		inOut := strings.Split(line, " -> ")
		rules[inOut[0]] = inOut[1]
	}
	return &rules
}

func templateToCounter(template string) *Counter {
	counter := make(Counter)
	for i := 0; i < len(template)-1; i++ {
		counter[template[i:i+2]] += 1
	}
	counter[" "+string(template[0])] = 1
	counter[string(template[len(template)-1])+" "] = 1
	return &counter
}

func (c *Counter) makeStep(rules *Rules) *Counter {
	c2 := make(Counter)
	var key2, key3 string
	for key, value := range *c {
		toInsert, found := (*rules)[key]
		if !found {
			c2[key] += value
			continue
		} else {
			key2 = string(key[0]) + toInsert
			key3 = toInsert + string(key[1])
		}
		c2[key2] += value
		c2[key3] += value
	}
	return &c2
}

func (c *Counter) count() *Counter {
	cnt := make(Counter)
	for key, value := range *c {
		if key[0] != ' ' {
			cnt[string(key[0])] += value
		}
	}
	return &cnt
}

func (c *Counter) minMax() (int, int) {
	min, max := math.MaxInt, 0
	for _, value := range *c {
		min = utils.MinInt(value, min)
		max = utils.MaxInt(value, max)
	}
	fmt.Println(min, max)
	return min, max
}
