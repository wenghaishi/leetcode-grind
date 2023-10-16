import java.util.ArrayList;
import java.util.List;


// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.

public class pascalTriangle {
  public static void main(String[] args) {
    System.out.println(getRow(0));
    System.out.println(getRow(1));
    System.out.println(getRow(2));
    System.out.println(getRow(3));
    System.out.println(getRow(4));

  }

  public static List<Integer> getRow(int rowIndex) {
    List<Integer> solution = new ArrayList<Integer>();
    solution.add(1);
    int currentRow = 0;
    while (currentRow != rowIndex) {
      currentRow++;
      solution.add(0, 0);
      solution.add(0);
      List<Integer> temp = new ArrayList<Integer>();
      for (int i = 0; i < solution.size() - 1; i++) {
        temp.add(solution.get(i) + solution.get(i + 1));
      }
      solution = temp;
    }
    return solution;
  }
}
