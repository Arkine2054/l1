package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, t := range data {
		key := int(t/10) * 10
		groups[key] = append(groups[key], t)

	}

	var keys []int
	for k := range groups {
		keys = append(keys, k)

	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}
