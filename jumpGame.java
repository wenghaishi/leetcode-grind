// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// difficulty: medium

public class jumpGame {
  public static void main(String[] args) {
    int[] nums = { 3, 2, 1, 0, 4 };
    System.out.println(jumpGameSolution(nums));
  }

  public static boolean jumpGameSolution(int[] nums) {
    int lastGoodIndexPosition = nums.length -1;
    for (int i = 0; i < nums.length -1; i--) {
      if ( nums[i] + i >= lastGoodIndexPosition) {
        lastGoodIndexPosition = i;
      } else {
        return false;
      }
    }
    return true;
  }
}
