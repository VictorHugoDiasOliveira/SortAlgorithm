package main

import (
	"fmt"

	sort "github.com/VictorHugoDiasOliveira/SortAlgorithm/golang/quick_sort"
	generator "github.com/VictorHugoDiasOliveira/UtilsAlgorithm/slicegenerator"
)

func main() {
	// usage of Quiksort algorithm
	array := generator.GenerateSliceWithRandomNumbers(40, 100)
	fmt.Println(array)
	sort.QuickSort(array[:], 0, len(array)-1)
	fmt.Println(array)
}
