package utils

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
 }

 func AbsDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }

 func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
 }


 func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
 }
