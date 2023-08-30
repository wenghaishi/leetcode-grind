// Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.

// A subarray is a contiguous part of an array.

// Example 1:

// Input: nums = [4,5,0,-2,-3,1], k = 5
// Output: 7
// Explanation: There are 7 subarrays with a sum divisible by k = 5:
// [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
// Example 2:

// Input: nums = [5], k = 9
// Output: 0

// difficulty: medium


// brute force approach
const k = 7;
const nums = [2, 5, 1, 3, 2, -2];
const subarrayK = (nums, k) => {
  let count = 0;
  let sum = 0;
  for (let i = 0; i < nums.length; i++) {
    sum = nums[0];
    for (let j = i + 1; j < nums.length; j++) {
      sum = nums[j];
      if (sum === k) count++;
    }
  }
  return count;
};
