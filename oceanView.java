import java.util.ArrayList;
import java.util.Arrays;

public class oceanView {
  public static void main(String[] args) {
    ArrayList<Integer> houses = new ArrayList<>(Arrays.asList(3, 7, 2, 8, 9));
    System.out.println(solution(houses));
  }

  public static int solution(ArrayList<Integer> houses) {
    int max = 0;
    int count = 0;
    for (int i = 0; i < houses.size(); i++) {
      if (houses.get(i) > max) {
        max = houses.get(i);
        count++;
      }
    }
    return count;
  }
}
