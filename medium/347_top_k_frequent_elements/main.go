package main

import "fmt"

type NumFreq struct {
	Num  int
	Freq int
}

func quickSort(a []NumFreq) []NumFreq {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := (right + 1) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Freq > a[right].Freq {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
func topKFrequent(nums []int, k int) []int {
	mappedFreqs := map[int]int{}

	for _, num := range nums {
		mappedFreqs[num]++
	}

	numFreqs := make([]NumFreq, 0, len(mappedFreqs))
	for num, freq := range mappedFreqs {
		numFreqs = append(numFreqs, NumFreq{
			Num:  num,
			Freq: freq,
		})
	}
	numFreqs = quickSort(numFreqs)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = numFreqs[i].Num
	}

	return res
}
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
