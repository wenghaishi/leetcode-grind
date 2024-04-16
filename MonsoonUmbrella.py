
def getUmbrellas(requirement, sizes):
    # Large number to represent infinity
    inf = float('inf')
    
    # Initialize dp array where dp[i] is the minimum umbrellas needed for i people
    dp = [inf] * (requirement + 1)
    dp[0] = 0  # 0 umbrellas needed to cover 0 people
    
    # Fill the dp array
    for size in sizes:
        for i in range(size, requirement + 1):
            if dp[i - size] != inf:
                dp[i] = min(dp[i], dp[i - size] + 1)
    
    # If dp[requirement] is still inf, it means it's not possible to cover exactly requirement people
    return dp[requirement] if dp[requirement] != inf else -1
