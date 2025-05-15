from collections import defaultdict

# 169. Majority Element

# Given an array nums of size n, return the majority element.

# The majority element is the element that appears more than ⌊n / 2⌋ times. 
# You may assume that the majority element always exists in the array.

def majority_element(nums):
    majority_size = len(nums) / 2
    frequencies = defaultdict(int)
    for num in nums:
        frequencies[num] += 1
        if frequencies[num] > majority_size:
            return num
        

print(majority_element([1,1,1,2,3]))