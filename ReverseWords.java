import java.util.ArrayList;

public class ReverseWords {
  public static void main(String[] args) {
    solution("hi    jack ");
  }

  public static String solution(String s) {
    String[] arr = s.trim().split(" ");
    System.out.println(arr);
    int upper = arr.length - 1;
    int lower = 0;
    while (upper > lower) {
      String temp = arr[lower].trim();
      arr[lower] = arr[upper].trim();
      arr[upper] = temp;
      upper--;
      lower++;
    }
    return String.join(" ", arr);
  }
}
