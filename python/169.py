def solution(nums):
    if len(nums) == 1:
        return nums[0]
    
    needed_frequencies = len(nums) // 2
    
    frequencies = {}

    for num in nums:
        frequencies[num] = frequencies.get(num, 0) + 1

        if frequencies[num] > needed_frequencies:
            return num

# Case 1: Simple majority element
print(solution([3, 3, 4, 2, 3, 3, 3]))  # Expected: 3

# Case 2: Majority at the beginning
print(solution([1, 1, 1, 2, 3, 4, 1]))  # Expected: 1

# Case 3: Only one element
print(solution([5]))  # Expected: 5