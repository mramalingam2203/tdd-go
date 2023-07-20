package main

// FindMax finds the maximum element in an array
func FindMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}
