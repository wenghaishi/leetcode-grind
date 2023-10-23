import java.util.ArrayList;
import java.util.List;

// There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

// Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

// Note that multiple kids can have the greatest number of candies.

public class KidsWithGreatestNumberOfCandies {
  public static void main(String[] args) {
    testSolution();
  }

  public static void testSolution() {
    // Test case 1: Regular case
    int[] candies1 = { 2, 3, 5, 1, 3 };
    int extraCandies1 = 3;
    List<Boolean> result1 = solution(candies1, extraCandies1);
    System.out.println("Test case 1: " + result1);

    // Test case 2: All children have the same number of candies
    int[] candies2 = { 2, 2, 2, 2 };
    int extraCandies2 = 1;
    List<Boolean> result2 = solution(candies2, extraCandies2);
    System.out.println("Test case 2: " + result2);

    // Test case 3: All children have different numbers of candies
    int[] candies3 = { 1, 2, 3, 4, 5 };
    int extraCandies3 = 2;
    List<Boolean> result3 = solution(candies3, extraCandies3);
    System.out.println("Test case 3: " + result3);

    // Test case 4: No extra candies
    int[] candies4 = { 1, 2, 3, 4, 5 };
    int extraCandies4 = 0;
    List<Boolean> result4 = solution(candies4, extraCandies4);
    System.out.println("Test case 4: " + result4);
  }

  public static List<Boolean> solution(int[] candies, int extraCandies) {
    int maxCandies = 0;
    ArrayList<Boolean> result = new ArrayList<>();
    for (int i = 0; i < candies.length; i++) {
      if (candies[i] > maxCandies)
        maxCandies = candies[i];
    }
    for (int j = 0; j < candies.length; j++) {
      if (candies[j] + extraCandies >= maxCandies) {
        result.add(true);
      } else {
        result.add(false);
      }
    }
    return result;

  }
}
