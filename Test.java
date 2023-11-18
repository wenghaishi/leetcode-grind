import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class Test {
  public static void main(String[] args) {
    Set<Integer> newSet = new HashSet<>();
    newSet.add(4);
    newSet.add(10);
    newSet.forEach((e) -> {
      System.out.println(e);
    });
  }

}
