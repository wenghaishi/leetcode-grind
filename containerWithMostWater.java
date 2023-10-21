// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.


public class containerWithMostWater {
  public static void main(String[] args) {
    int[] arr = {1, 8, 6, 2, 5, 4, 8, 3, 7};
    System.out.println(maxArea(arr));
  }

  public static int maxArea(int[] height) {
    int lowerIndex = 0;
    int upperIndex = height.length - 1;
    int maxArea = 0;
    while (upperIndex > lowerIndex) {
        int currentHeight = Math.min(height[upperIndex], height[lowerIndex]);
        int currentWidth = upperIndex - lowerIndex;
        int currentArea = currentHeight * currentWidth;
        if (currentArea > maxArea) {
            maxArea = currentArea;
        }
        if (height[upperIndex] > height[lowerIndex]) {
            lowerIndex++;
        } else {
            upperIndex--;
        }
    }
    return maxArea;
  }
}
