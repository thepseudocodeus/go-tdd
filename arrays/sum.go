package arrays

/*
 * Maybe new requirement:
 *
 * Convert the input slice to a fixed size array.
 *
 * Why? Interesting problem to solve.
 */

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
