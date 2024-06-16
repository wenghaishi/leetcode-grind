package sorting;

import java.util.ArrayList;
import java.util.Arrays;

public class InsertionSort {
    public static void main(String args[]) {
        // Test case 1: Regular case
        ArrayList<Integer> list1 = new ArrayList<>(Arrays.asList(5, 2, 9, 1, 5, 6));
        System.out.println("Test case 1: " + insertionSort(list1)); // Expected: [1, 2, 5, 5, 6, 9]

        // Test case 2: Empty list
        ArrayList<Integer> list2 = new ArrayList<>();
        System.out.println("Test case 2: " + insertionSort(list2)); // Expected: []

        // Test case 3: List with one element
        ArrayList<Integer> list3 = new ArrayList<>(Arrays.asList(1));
        System.out.println("Test case 3: " + insertionSort(list3)); // Expected: [1]

        // Test case 4: List with duplicate elements
        ArrayList<Integer> list4 = new ArrayList<>(Arrays.asList(3, 3, 3));
        System.out.println("Test case 4: " + insertionSort(list4)); // Expected: [3, 3, 3]

        // Test case 5: List already sorted
        ArrayList<Integer> list5 = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        System.out.println("Test case 5: " + insertionSort(list5)); // Expected: [1, 2, 3, 4, 5]

        // Test case 6: List sorted in reverse order
        ArrayList<Integer> list6 = new ArrayList<>(Arrays.asList(5, 4, 3, 2, 1));
        System.out.println("Test case 6: " + insertionSort(list6)); // Expected: [1, 2, 3, 4, 5]
    }

    public static ArrayList<Integer> insertionSort(ArrayList<Integer> input) {
        ArrayList<Integer> solution = new ArrayList<Integer>();
        for (int i = 0; i < input.size(); i++) {
            Integer current = input.get(i);

            if (solution.size() == 0) {
                solution.add(current);
            } else {
                boolean inserted = false;
                for (int j = 0; j < solution.size(); j++) {
                    if (current < solution.get(j)) {
                        solution.add(j, current);
                        inserted = true;
                        break;
                    }
                }
                if (!inserted) {
                    solution.add(current);
                }
            }
        }
        return solution;
    }
}
