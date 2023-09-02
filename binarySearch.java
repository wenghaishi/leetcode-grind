public class binarySearch {
  public static void main(String[] args) {
    int[] nums = { -1, 0, 3, 5, 9, 12 };
    int target = 0;
    System.out.println(bSearch(nums, target));
  }

  public static int bSearch(int[] nums, int target) {
    int lowerIndex = 0;
    int upperIndex = nums.length - 1;
    
    while (upperIndex >= lowerIndex) {
        int middleIndex = lowerIndex + (upperIndex - lowerIndex) / 2; // Calculate the middle index correctly.

        if (target == nums[middleIndex])
            return middleIndex;
        if (target > nums[middleIndex]) {
            lowerIndex = middleIndex + 1;
        } else {
            upperIndex = middleIndex - 1;
        }
    }
    return -1;
  }
}
