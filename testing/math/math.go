//Package math provides a set of math utility functions
package math

//MySum adds the passed series of ints to return the total
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
