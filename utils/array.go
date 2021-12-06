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
