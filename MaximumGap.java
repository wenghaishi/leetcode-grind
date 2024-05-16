public class MaximumGap {
    
    public static void main(String[] args) {

    }

    public static int solution(int[] nums) {
        if (nums.length < 2) return 0;

        int max = 0;

        for (int i = 0; i < nums.length; i++) {
            if (i != nums.length -1) {
                int diff = nums[i + 1] - nums[i];
                if (diff > max) {
                    max = diff;
                }
            }
        }

        return max;
    }
}
