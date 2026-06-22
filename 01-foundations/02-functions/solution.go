package functions

import "fmt"

// Sum returns the total of all integers passed in. Sum() with no arguments returns 0.
func Sum(nums ...int) int {
	total:=0
	for _,n := range nums{
		total+=n
	}
	return total
}

// Divide returns integer quotient a/b. When b is zero, return an error instead of panicking.
func Divide(a, b int) (int, error) {
	if(b==0){
		return 0, fmt.Errorf("Can't divide by zero")
	}
	return a/b, nil
}
