import java.util.HashSet;

public class MissingNum {
    public static void main(String[] args) {

    }

    public static int missingNumber(int[] nums) {
        HashSet<Integer> numbers = new HashSet<>();

        // Add all numbers in the array to the HashSet
        for (int num : nums) {
            numbers.add(num);
        }

        // Check for the missing number in the range [0, nums.length]
        for (int i = 0; i <= nums.length; i++) {
            if (!numbers.contains(i)) {
                return i;
            }
        }

        return 0; // Should not reach here if input is valid
    }
}
