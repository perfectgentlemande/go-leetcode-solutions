package main

import "fmt"

func average(salary []int) float64 {
	min, max := 0, 0

	for i := 1; i < len(salary); i++ {
		if salary[i] > salary[max] {
			max = i
		}
		if salary[i] < salary[min] {
			min = i
		}
	}

	sum := 0
	for i := range salary {
		if i != max && i != min {
			sum += salary[i]
		}
	}

	return float64(sum) / float64(len(salary)-2)
}

func main() {
	fmt.Println(average([]int{4000, 3000, 1000, 2000}))
	fmt.Println(average([]int{1000, 2000, 3000}))
}
