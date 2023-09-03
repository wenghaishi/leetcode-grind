// given a number, add the first and last digit of it

// difficulty: easy

public class firstLastDigitSum {
  public static void main(String[] args) {
    System.out.println(sum(11));
  }

  public static int sum(int number) {
    String numberAsString = Integer.toString(number);
    int firstNumber = Character.getNumericValue(numberAsString.charAt(0));
    int lastNumber = Character.getNumericValue(numberAsString.charAt(numberAsString.length() - 1));

    return firstNumber + lastNumber;
  }
}
