import java.util.ArrayList;

public class PlusOne {
    public static void main(String args[]) {

    }

    public int[] plusOne(int[] digits) {
        ArrayList<Integer> solution = new ArrayList<>();

        int carry = 1;
        for (int i = digits.length - 1; i >= 0; i--) {
            int sum = digits[i] + carry;
            if (sum > 9) {
                solution.add(0, 0);
            } else {
                solution.add(0, sum);
                carry = 0;
            }
        }

        // If there's still a carry left after the last addition, add it to the front
        if (carry != 0) {
            solution.add(0, carry);
        }

        // Convert ArrayList<Integer> to int[]
        int[] solutionArray = new int[solution.size()];
        for (int i = 0; i < solution.size(); i++) {
            solutionArray[i] = solution.get(i);
        }

        return solutionArray;
    }
}
