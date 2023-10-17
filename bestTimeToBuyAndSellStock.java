
// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

public class bestTimeToBuyAndSellStock {
  public static void main(String[] args) {
    int[] prices = { 2, 4, 2, 8, 2, 1 };
    System.out.println(solution(prices));
  }

  public static int solution(int[] prices) {
    int maxProfit = 0;
    for (int i = 0; i < prices.length; i++) {
      for (int j = i + 1; j < prices.length; j++) {
        int profit = prices[j] - prices[i];
        if (profit > maxProfit) {
          maxProfit = profit;
        }
      }
    }
    return maxProfit;
  }

  public static int betterSolution(int[] prices) {
    int maxProfit = 0;
    int upperIndex = 1;
    int lowerIndex = 0;
    while (upperIndex < prices.length) {
      if (prices[upperIndex] < prices[lowerIndex]) {
        lowerIndex = upperIndex;
        upperIndex++;
      } else {
        int profit = prices[upperIndex] - prices[lowerIndex];
        if (profit > maxProfit) {
          maxProfit = profit;
        }
        upperIndex++;
      }
    }
    return maxProfit;
  }

}
