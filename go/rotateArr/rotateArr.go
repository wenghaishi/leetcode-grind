package main

func rotateArr(nums []int, k int) {
	numsLength := len(nums)
	remainder := k % numsLength 
	newArr := make([]int, 0, numsLength)

	newArr = append(newArr, nums[numsLength - remainder:]...)
	newArr = append(newArr, nums[:numsLength - remainder]...)

	for i :=0; i < numsLength; i++ {
		nums[i] = newArr[i]
	}
}