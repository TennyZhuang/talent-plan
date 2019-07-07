package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	const maxInt64 = 9223372036854775807

	n := len(src)
	if n == 1 {
		return
	}
	m := n / 2
	MergeSort(src[:m])
	arr1 := make([]int64, m+1)
	copy(arr1, src[:m])
	arr1[m] = maxInt64

	MergeSort(src[m:])
	arr2 := make([]int64, n-m+1)
	copy(arr2, src[m:])
	arr2[n-m] = maxInt64

	i, j := 0, 0
	for k := 0; k < n; k++ {
		if arr1[i] <= arr2[j] {
			src[k] = arr1[i]
			i++
		} else {
			src[k] = arr2[j]
			j++
		}
	}
}
