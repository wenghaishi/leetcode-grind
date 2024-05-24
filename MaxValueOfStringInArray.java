public class MaxValueOfStringInArray {
    public static void main(String[] args) {

    }

    public int solution(String[] strs) {
        int maxValue = 0;

        for (int i = 0; i < strs.length; i++) {
            int value;
            if (strs[i].matches("\\d+")) {
                value = Integer.parseInt(strs[i]);
            } else {
                value = strs[i].length();
            }

            if (value > maxValue) {
                maxValue = value;
            }
        }

        return maxValue;
    }
}
