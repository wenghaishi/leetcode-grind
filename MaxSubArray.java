public class MaxSubArray {
    public static void main(String[] args) {
        testSingleElementArray();
        testAllPositiveNumbers();
        testAllNegativeNumbers();
        testMixedNumbers();
        testArrayWithZeros();
    }

    public static int solution(int[] nums) {
        int max = nums[0];
        int curr = nums[0];

        for (int i = 1; i < nums.length; i++) {
            curr = Math.max(nums[i], curr + nums[i]);

            if (curr > max) {
                max = curr;
            }
        }

        return max;
    }

    // Test cases

    public static void testSingleElementArray() {
        int[] nums = {5};
        int expected = 5;
        int result = solution(nums);
        System.out.println("Test Single Element Array: " + (result == expected));
    }

    public static void testAllPositiveNumbers() {
        int[] nums = {1, 2, 3, 4, 5};
        int expected = 15;
        int result = solution(nums);
        System.out.println("Test All Positive Numbers: " + (result == expected));
    }

    public static void testAllNegativeNumbers() {
        int[] nums = {-1, -2, -3, -4, -5};
        int expected = -1;
        int result = solution(nums);
        System.out.println("Test All Negative Numbers: " + (result == expected));
    }

    public static void testMixedNumbers() {
        int[] nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
        int expected = 6;
        int result = solution(nums);
        System.out.println("Test Mixed Numbers: " + (result == expected));
    }

    public static void testArrayWithZeros() {
        int[] nums = {0, -3, 1, 1, -1, 2, -1, 3, -5, 4};
        int expected = 5;
        int result = solution(nums);
        System.out.println("Test Array with Zeros: " + (result == expected));
    }
}
