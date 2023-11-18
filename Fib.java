public class Fib {
  public static void main(String[] args) {
    System.out.println(solution(10));
    System.out.println(solution(6));
  }

  public static int solution(int num) {
    if (num == 0) return 0;
    if (num == 1) return 1;
    return solution(num -1) + solution(num - 2);
  }
}
