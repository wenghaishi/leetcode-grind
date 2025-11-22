package main

func mergeSort(arr []int) []int {

	inputLength :=len(arr)

	if inputLength < 2 {
		return
	}

	mid := inputLength / 2
	leftHalf := arr[:mid]
	rightHalf := arr[mid:]
}
