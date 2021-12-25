package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToInts(row string, split string) []int {
	var inList []string
	if split == " " {
		inList = strings.Fields(row)
	} else {
		inList = strings.Split(row, split)
	}
	nums := make([]int, len(inList))
	for i, s := range inList {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}

func PrintArray(array [][]int) {
	for _, row := range array {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func Sum(array []int) int {
	result := 0
	for _, el := range array {
		result += el
	}
	return result
}

func ReverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
