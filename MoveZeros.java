import java.util.Arrays;

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

public class MoveZeros {
  public static void main(String[] args) {

    // Test Case 1: Basic test case with non-zero and zero elements
    int[] nums1 = { 0, 1, 0, 3, 12 };
    solution(nums1);
    System.out.println(Arrays.toString(nums1));

    // Test Case 2: All elements are non-zero
    int[] nums2 = { 1, 2, 3, 4, 5 };
    solution(nums2);
    System.out.println(Arrays.toString(nums2));

    // Test Case 3: All elements are zeros
    int[] nums3 = { 0, 0, 0, 0, 0 };
    solution(nums3);
    System.out.println(Arrays.toString(nums3));

  }

  public static void solution(int[] nums) {
    if (nums.length == 0 || nums == null)
      return;
    int moveTo = 0;
    for (int num : nums) {
      if (num != 0) {
        nums[moveTo++] = num;
      }
    }

    while (moveTo < nums.length) {
      nums[moveTo++] = 0;
    }
  }
}
