public class BestTimeToBuyAndSellStockTwo {
  public static void main(String[] args) {
    int[] prices1 = { 1, 2, 3 };
    int[] prices2 = { 7, 1, 5, 3, 6, 4 };
    int[] prices3 = { 5, 1 };

    System.out.println(solution(prices1));
    System.out.println(solution(prices2));
    System.out.println(solution(prices3));

  }

  public static int solution(int[] prices) {
    int profit = 0;
    for (int i = 1; i < prices.length; i++) {
      if (prices[i] - prices[i - 1] > 0) {
        profit += prices[i] - prices[i - 1];
      }
    }
    return profit;
  }
}


public static void goodbye() {
  systems.out.println()
}
