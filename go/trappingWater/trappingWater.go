package main

func trappingWater(heights []int) int {
    n := len(heights)
    if n == 0 {
        return 0
    }

    maxLeft := make([]int, n)
    maxRight := make([]int, n)

    // prefix max
    maxLeft[0] = heights[0]
    for i := 1; i < n; i++ {
        maxLeft[i] = max(maxLeft[i-1], heights[i])
    }

    // suffix max
    maxRight[n-1] = heights[n-1]
    for i := n - 2; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], heights[i])
    }

    total := 0
    for i := 0; i < n; i++ {
        water := min(maxLeft[i], maxRight[i]) - heights[i]
        if water > 0 {
            total += water
        }
    }

    return total

}
