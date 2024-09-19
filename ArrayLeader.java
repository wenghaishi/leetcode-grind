import java.util.ArrayList;

public class ArrayLeader {
    public static void main(String[] args) {
        int[] arr1 = { 1, 3, 6, 1 };

        int[] arr1Solution = solution(arr1);

        for (int i = 0; i < arr1Solution.length; i++) {
            System.out.println(arr1Solution[i]);
        }

    }

    public static int[] solution(int[] arr) {
        int currentMax = arr[arr.length - 1];

        ArrayList<Integer> solutionArrayList = new ArrayList<Integer>();

        for (int i = arr.length - 1; i >= 0; i--) {
            if (arr[i] >= currentMax) {
                solutionArrayList.add(Integer.valueOf(arr[i]));
                currentMax = arr[i];
            }
        }

        int[] solutionArr = solutionArrayList.stream().mapToInt(i -> i).toArray();

        return solutionArr;

    }
}
