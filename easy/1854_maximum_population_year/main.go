package main

import (
	"sort"
)

func maximumPopulation(logs [][]int) int {
	mp := map[int]int{}
	years := []int{}

	for i := range logs {
		if _, ok := mp[logs[i][0]]; !ok {
			years = append(years, logs[i][0])
		}

		if _, ok := mp[logs[i][1]]; !ok {
			years = append(years, logs[i][1])
		}

		mp[logs[i][0]]++
		mp[logs[i][1]]--
	}

	sort.Ints(years)

	maxCount := 0
	currCount := 0
	maxYear := years[0]

	for _, year := range years {
		currCount += mp[year]

		if currCount > maxCount {
			maxCount = currCount
			maxYear = year
		}
	}

	return maxYear
}
