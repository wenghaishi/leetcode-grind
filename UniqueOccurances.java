import java.util.HashSet;

public class UniqueOccurances {
    public static void main(String[] args) {

    }

    public static boolean solution(int[] input) {
        HashSet<Integer> mySet = new HashSet<>();
        for (int i = 0; i < input.length; i++) {
            if (mySet.contains(Integer.valueOf(input[i]))) {
                return false;
            } else {
                mySet.add(Integer.valueOf(input[i]));
            }
        }

        return true;
    }
}
