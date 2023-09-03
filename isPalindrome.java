// isPalindrome?
// return true if a number is palinedrome, otherwise false

// difficulty: easy


public class isPalindrome {
  public static void main(String[] args) {
    System.out.println(isNumberPalindrome(12221));
  }

  public static boolean isNumberPalindrome(int number) {
    String numberAsString = Integer.toString(number);
    int length = numberAsString.length();
    for (int i = 0; i < length / 2; i++) {
      if (numberAsString.charAt(i) != numberAsString.charAt(length - 1 - i))
        return false;
    }
    return true;
  }
}
