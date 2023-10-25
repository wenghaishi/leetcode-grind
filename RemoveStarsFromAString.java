// You are given a string s, which contains stars *.

// In one operation, you can:

// Choose a star in s.
// Remove the closest non-star character to its left, as well as remove the star itself.
// Return the string after all stars have been removed.

public class RemoveStarsFromAString {
  public static void main(String[] args) {
    System.out.println(solution("leetcode****"));
  }

  public static String solution(String s) {
    StringBuilder newString = new StringBuilder();
    for (int i = 0; i < s.length(); i++) {
      if (s.charAt(i) == '*') {
        newString.deleteCharAt(newString.length() - 1);
      } else {
        newString.append(s.charAt(i));
      }
    }
    String result = newString.toString();
    return result;
  }
}
