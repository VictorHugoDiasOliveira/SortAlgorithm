package main

import (
	"fmt"

	sort "github.com/VictorHugoDiasOliveira/SortAlgorithm/golang/quick_sort"
)

func main() {
	// usage of Quiksort algorithm
	array := [8]int{4, 2, 8, 7, 1, 5, 3, 6}
	fmt.Println(array)
	sort.QuickSort(array[:], 0, len(array)-1)
	fmt.Println(array)
}
