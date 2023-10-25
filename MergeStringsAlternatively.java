// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, 
//starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

public class MergeStringsAlternatively {
  public static void main(String[] args) {
    System.out.println(solution("ace", "bd"));
  }

  public static String solution(String word1, String word2) {
    StringBuilder newStr = new StringBuilder();
    for (int i = 0; i < Math.max(word1.length(), word2.length()); i++) {
      if (i < word1.length()) {
        newStr.append(word1.charAt(i));
      }
      if (i < word2.length()) {
        newStr.append(word2.charAt(i));
      }
    }
    String solution = newStr.toString();
    return solution;
  }
}
