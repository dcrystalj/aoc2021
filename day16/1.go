package day16

import (
	"fmt"
	"math"
	"strconv"

	"github.com/dcrystalj/aoc2021/registrator"
	"github.com/dcrystalj/aoc2021/utils"
)

var versionSum int

func init() {
	registrator.Register("day16.1", Solve)
	registrator.Register("day16.2", Solve2)

}

func Solve(lines []string) string {
	input := lineToBin(lines[0])
	parsePacket(input)
	return strconv.Itoa(versionSum)
}

func Solve2(lines []string) string {
	input := lineToBin(lines[0])
	tree, _ := parsePacket(input)
	res := evalTree(tree)
	return fmt.Sprintf("%d", res)
}

type Literal struct {
	version int
	id      int
	value   int64
}

type Operator struct {
	version int
	id      int
	packets []interface{}
}

func lineToBin(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		parsedInt, _ := strconv.ParseInt(s[i:i+1], 16, 32)
		res += fmt.Sprintf("%04b", parsedInt)
	}
	return res
}

func parsePacket(s string) (interface{}, string) {
	if len(s) < 6 {
		return nil, s
	}
	id := s[3:6]
	if id == "100" {
		return NewLiteral(s)
	} else {
		return NewOperator(s)
	}
}

func NewLiteral(s string) (*Literal, string) {
	version1, _ := strconv.ParseInt(s[:3], 2, 32)
	version := int(version1)
	versionSum += version
	s = s[6:]
	toBuild := ""
	remaining := ""
	for {
		toBuild += s[1:5]
		if s[0] == '0' {
			remaining = s[5:]
			break
		} else {
			s = s[5:]
		}
	}
	value, _ := strconv.ParseInt(toBuild, 2, 64)

	return &Literal{
		version, 4, int64(value),
	}, remaining
}

func NewOperator(s string) (o *Operator, remaining string) {
	version1, _ := strconv.ParseInt(s[:3], 2, 32)
	version := int(version1)
	versionSum += version
	s = s[3:]
	id1, _ := strconv.ParseInt(s[:3], 2, 32)
	id := int(id1)
	s = s[3:]
	packets := make([]interface{}, 0)
	if s[0] == '0' {
		s = s[1:]
		subPacketsLen1, _ := strconv.ParseInt(s[:15], 2, 32)
		subPacketsLen := int(subPacketsLen1)
		s = s[15:]
		toParse := s[:subPacketsLen]
		for {
			packet, r := parsePacket(toParse)
			toParse = r
			if packet == nil {
				break
			}
			packets = append(packets, packet)
		}
		remaining = s[subPacketsLen:]
	} else {
		s = s[1:]
		nPackets1, _ := strconv.ParseInt(s[:11], 2, 32)
		nPackets := int(nPackets1)
		toParse := s[11:]
		for i := 0; i < nPackets; i++ {
			packet, r := parsePacket(toParse)
			packets = append(packets, packet)
			remaining = r
			toParse = r
		}
	}
	return &Operator{
		version, id, packets,
	}, remaining
}

func evalTree(tree interface{}) int64 {
	switch t := tree.(type) {
	case nil:
		fmt.Print("Nil")
		return 0
	case *Literal:
		return int64(t.value)
	case *Operator:
		switch t.id {
		case 0:
			res := int64(0)
			for _, packet := range t.packets {
				res += evalTree(packet)
			}
			return res
		case 1:
			res := int64(1)
			for _, packet := range t.packets {
				res *= evalTree(packet)
			}
			return res
		case 2:
			res := int64(math.MaxInt64)
			for _, packet := range t.packets {
				res = utils.MinInt64(res, evalTree(packet))
			}
			return res
		case 3:
			res := int64(math.MinInt64)
			for _, packet := range t.packets {
				res = utils.MaxInt64(res, evalTree(packet))
			}
			return res
		case 5:
			first := evalTree(t.packets[0])
			second := evalTree(t.packets[1])
			if first > second {
				return 1
			}
			return 0
		case 6:
			first := evalTree(t.packets[0])
			second := evalTree(t.packets[1])
			if first < second {
				return 1
			}
			return 0
		case 7:
			first := evalTree(t.packets[0])
			second := evalTree(t.packets[1])
			if first == second {
				return 1
			}
			return 0
		}
	}
	fmt.Print("Not implme")
	return 0
}
