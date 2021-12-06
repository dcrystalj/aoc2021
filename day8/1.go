package day8

import (
	"strconv"
	"strings"

	"github.com/dcrystalj/aoc2021/registrator"
	"k8s.io/apimachinery/pkg/util/sets"
)


func init() {
    registrator.Register("day8.1", Solve)
    registrator.Register("day8.2", Solve2)
}


type Line struct {
	in []string
	out []string
}

func readLines(lines []string) []Line {
	var rows []Line
	for _, line := range lines {
		firstSplit := strings.Split(line, " | ")
		input := firstSplit[0]
		output := firstSplit[1]
		rows = append(rows, Line{
			strings.Split(input, " "),
			strings.Split(output, " "),
		})
	}
	return rows
}

func lenInMapper(line []string) map[int][]sets.Byte {
	mapper := make(map[int][]sets.Byte, 0)
	for _, num := range line {
		set := sets.NewByte()
		for i := 0; i < len(num); i++ {
			set.Insert(num[i])
		}

		_, ok := mapper[len(num)]
		if !ok {
			mapper[len(num)] = make([]sets.Byte, 0)
		}
		mapper[len(num)] = append(mapper[len(num)], set)

	}
	return mapper
}

func getVal(mapper map[int][]sets.Byte, key int) []sets.Byte {
	val := mapper[key]
	return val
}

func mergeSets(sets []sets.Byte) sets.Byte{
	firsSet := sets[0]
	for _, set := range sets {
		firsSet = firsSet.Intersection(set)
	}
	return firsSet
}

func lenMapperToValue(lenMapper map[int][]sets.Byte) map[byte]byte {
	options := make(map[byte]sets.Byte)

	len2 := mergeSets(getVal(lenMapper, 2))
	len3 := mergeSets(getVal(lenMapper, 3))
	len4 := mergeSets(getVal(lenMapper, 4))
	len5 := mergeSets(getVal(lenMapper, 5))
	len6 := mergeSets(getVal(lenMapper, 6))
	len7 := mergeSets(getVal(lenMapper, 7))

	options['a']=(len3.Intersection(len5).Intersection(len6).Intersection(len7)).Difference(len2).Difference(len4)
	options['b']=(len4.Intersection(len6).Intersection(len7)).Difference(len2).Difference(len3)
	options['c']=(len2.Intersection(len3).Intersection(len4).Intersection(len7))
	options['d']=(len4.Intersection(len5).Intersection(len7)).Difference(len2).Difference(len3)
	options['e']=(len7).Difference(len2).Difference(len3).Difference(len4)
	options['f']=(len2.Intersection(len3).Intersection(len4).Intersection(len6).Intersection(len7))
	options['g']=(len5.Intersection(len6).Intersection(len7)).Difference(len2).Difference(len3).Difference(len4)

	options['c'] = options['c'].Difference(options['f'])
	options['e'] = options['e'].Difference(options['g'])

	return optimizeOpions(options)
}

func optimizeOpions(options map[byte]sets.Byte) map[byte]byte {
	result := make(map[byte]byte)
	for key, val := range options {
		if val.Len() != 1 {
			panic("Not optimized")
		}
		if (val.Len() == 1) {
			result[key] = val.List()[0]
		}
	}
	return result
}

func convertOuputLine(options map[byte]byte, outs []string) int { // strings together
	var s []byte
	for _, out := range outs {
		s = append(s, convertWordToNum(options, out))
	}
	i, _ := strconv.ParseInt(string(s), 10, 32)
	return int(i)
}

func mapSet(options map[byte]byte, set sets.Byte) sets.Byte {
	result := sets.NewByte()
	for s := range set {
		result.Insert(options[s])
	}
	return result
}

func convertWordToNum(options map[byte]byte, out string) byte { //get one num
	mapped := sets.NewByte()
	for _, char := range out {
		mapped.Insert(byte(char))
	}

	switch {
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'b', 'c', 'e', 'f', 'g'))):
		return '0'
	case mapped.Equal(mapSet(options, sets.NewByte('c', 'f'))):
		return '1'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'c', 'd', 'e', 'g'))):
		return '2'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'c', 'd', 'f', 'g'))):
		return '3'
	case mapped.Equal(mapSet(options, sets.NewByte('b', 'c', 'd', 'f'))):
		return '4'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'b', 'd', 'f', 'g'))):
		return '5'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'b', 'd', 'e', 'f', 'g'))):
		return '6'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'c', 'f'))):
		return '7'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'b', 'c', 'd', 'e', 'f', 'g'))):
		return '8'
	case mapped.Equal(mapSet(options, sets.NewByte('a', 'b', 'c', 'd', 'f', 'g'))):
		return '9'
	}
	panic("Invalid")
}

func Solve(inLines []string) string {
	cnt := 0
	lines := readLines(inLines)
	for _, line := range lines {
		for _, out := range line.out {
			if len(out) == 2 || len(out) == 3 || len(out) == 4 || len(out) == 7 {
				cnt ++
			}
		}
	}

	return strconv.Itoa(cnt)
}

func Solve2(inLines []string) string {
	cnt := 0
	lines := readLines(inLines)
	for _, line := range lines {
		mapper := lenInMapper(line.in)
		options := lenMapperToValue(mapper)
		cnt += convertOuputLine(options, line.out)
	}
	return strconv.Itoa(cnt)
}
