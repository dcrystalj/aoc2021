package day18

import (
	"testing"
)

var addStacksParams = []struct {
	in1, in2, expected string
}{
	{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[[4 3 ]4 ]4 ][7 [[8 4 ]9 ]]][1 1 ]]"},
}

func TestAddStacks(t *testing.T) {
	for _, tt := range addStacksParams {
		t.Run("t", func(t *testing.T) {
			s1 := createStack(tt.in1)
			s2 := createStack(tt.in2)
			result := serializeStack(addStacks(s1, s2))
			if result != tt.expected {
				t.Errorf("%s Not equal %s", result, tt.expected)
			}
		})
	}
}

var sumLinesParams = []struct {
	lines    []string
	expected string
}{
	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}, "[[[[1 1 ][2 2 ]][3 3 ]][4 4 ]]"},
	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}, "[[[[3 0 ][5 3 ]][4 4 ]][5 5 ]]"},
	{[]string{"[1,9]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}, "[[[[0 6 ][5 3 ]][4 4 ]][5 5 ]]"},
	{
		[]string{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			"[7,[5,[[3,8],[1,4]]]]",
			"[[2,[2,2]],[8,[8,1]]]",
			"[2,9]",
			"[1,[[[9,3],9],[[9,0],[0,7]]]]",
			"[[[5,[7,4]],7],1]",
			"[[[[4,2],2],6],[8,7]]",
		},
		"[[[[8 7 ][7 7 ]][[8 6 ][7 7 ]]][[[0 7 ][6 6 ]][8 7 ]]]",
	},
	{
		[]string{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		},
		"[[[[4 0 ][5 4 ]][[7 7 ][6 0 ]]][[8 [7 7 ]][[7 9 ][5 0 ]]]]",
	},
}

func TestSumLines(t *testing.T) {
	for _, tt := range sumLinesParams {
		t.Run("t", func(t *testing.T) {
			result := serializeStack(sumLines(tt.lines))
			if result != tt.expected {
				t.Errorf("%s Not equal %s", result, tt.expected)
			}
		})
	}
}

var explodeParams = []struct {
	line            string
	explodePosition int
	expected        string
}{
	{"[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]", 4, "[[[[0 [3 2 ]][3 3 ]][4 4 ]][5 5 ]]"},
	{"[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]", 5, "[[[[3 0 ][5 3 ]][4 4 ]][5 5 ]]"},
}

func TestExplode(t *testing.T) {
	for _, tt := range explodeParams {
		t.Run("t", func(t *testing.T) {
			s := createStack(tt.line)
			result := serializeStack(explode(s, tt.explodePosition))
			if result != tt.expected {
				t.Errorf("%s Not equal %s", result, tt.expected)
			}
		})
	}
}

// var splitParams = []struct {
// 	line          string
// 	splitPosition int
// 	expected      string
// }{
// 	{"[[[[11,1],[2,2]],[3,3]],[4,4]]", 4, "[[[[[5 6 ] 1] [3 2 ]][3 3 ]][4 4 ]][5 5 ]]"},
// }

// func TestSplit(t *testing.T) {
// 	for _, tt := range splitParams {
// 		t.Run("t", func(t *testing.T) {
// 			s := createStack(tt.line)
// 			result := serializeStack(split(s, tt.splitPosition))
// 			if result != tt.expected {
// 				t.Errorf("%s Not equal %s", result, tt.expected)
// 			}
// 		})
// 	}
// }

var optimizeParams = []struct {
	line     string
	expected string
}{
	{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0 7 ]4 ][[7 8 ][6 0 ]]][8 1 ]]"},
}

func TestOptimize(t *testing.T) {
	for _, tt := range optimizeParams {
		t.Run("t", func(t *testing.T) {
			s := createStack(tt.line)
			result := serializeStack(optimize(s))
			if result != tt.expected {
				t.Errorf("%s Not equal %s", result, tt.expected)
			}
		})
	}
}
