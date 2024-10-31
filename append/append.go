package append

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	amountOfDays := 0
	for _, cost := range costs {
		if cost.day > amountOfDays {
			amountOfDays = cost.day
		}
	}

	result := make([]float64, amountOfDays+1)
	fmt.Println(result)

	for _, cost := range costs {
		result[cost.day] += cost.value
	}
	fmt.Println(result)
	return nil
}
