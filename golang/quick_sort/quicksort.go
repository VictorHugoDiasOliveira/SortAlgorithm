package quicksort

func QuickSort(arr []int, low int, high int) {

	if low < high {
		pivot := Partion(arr, low, high)

		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}

}

func Partion(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
