package main

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(arrays ...[]int) []int {

	var sums []int

	for _, numbers := range arrays {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrays ...[]int) []int {

	var sums []int

	for _, numbers := range arrays {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
