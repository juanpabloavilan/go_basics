package slices

func Sum(numbers []int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}

	return sums
}
