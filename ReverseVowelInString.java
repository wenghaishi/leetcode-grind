public class ReverseVowelInString {
  public static void main(String[] args) {

  }

  public String solution(String s) {
    String vowels = "aeiouAEIOU";
    char[] chars = s.toCharArray();
    int left = 0;
    int right = s.length() - 1;

    while (left < right) {
        while (left < right && vowels.indexOf(chars[left]) == -1) {
            left++;
        }
        while (left < right && vowels.indexOf(chars[right]) == -1) {
            right--;
        }

        if (left < right) {
            char temp = chars[left];
            chars[left] = chars[right];
            chars[right] = temp;
            left++;
            right--;
        }
    }

    return new String(chars);
}
}
