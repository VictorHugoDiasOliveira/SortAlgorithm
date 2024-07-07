package quicksort

func QuickSort(arr []int, start_position int, end_position int) {

	if start_position < end_position {
		pivot := Partion(arr, start_position, end_position)

		QuickSort(arr, start_position, pivot-1)
		QuickSort(arr, pivot+1, end_position)
	}

}

func Partion(arr []int, start_position int, end_position int) int {
	pivot := arr[end_position]
	i := start_position - 1

	for j := start_position; j < end_position; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end_position] = arr[end_position], arr[i+1]
	return i + 1
}
