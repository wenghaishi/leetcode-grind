

def solution(nums, target):
    upper = len(nums)
    lower = 0

    while upper >= lower:
        mid = (upper + lower) // 2

        if nums[mid] == target:
            prev = mid - 1
            next = mid + 1

            while (prev >= 0 and nums[prev] == target) or (next < len(nums) and nums[next] == target):
                if prev >= 0 and nums[prev] == target:
                    prev -= 1
                if next < len(nums) and nums[next] == target:
                    next += 1
            
            return [prev + 1, next - 1]
        
        elif nums[mid] > target:
            upper = mid - 1
        else:
            lower = mid + 1
    
    return [-1, -1]

def test_solution():
    # Test Case 1: Target exists with multiple occurrences
    assert solution([5, 7, 7, 8, 8, 10], 8) == [3, 4], "Test Case 1 Failed"

    # Test Case 2: Target exists once
    assert solution([5, 7, 7, 8, 10], 7) == [1, 2], "Test Case 2 Failed"

    # Test Case 3: Target not in array
    assert solution([1, 2, 3, 4, 5], 6) == [-1, -1], "Test Case 3 Failed"

    # Test Case 4: Empty array
    assert solution([], 1) == [-1, -1], "Test Case 4 Failed"

    # Test Case 5: Single element, target present
    assert solution([1], 1) == [0, 0], "Test Case 5 Failed"

    # Test Case 6: Single element, target absent
    assert solution([1], 2) == [-1, -1], "Test Case 6 Failed"

    # Test Case 7: Target at start
    assert solution([1, 1, 2, 3, 4], 1) == [0, 1], "Test Case 7 Failed"

    # Test Case 8: Target at end
    assert solution([1, 2, 3, 4, 4], 4) == [3, 4], "Test Case 8 Failed"

    # Test Case 9: All elements are target
    assert solution([2, 2, 2, 2], 2) == [0, 3], "Test Case 9 Failed"

    print("All test cases passed!")

# Run the tests
try:
    test_solution()
except AssertionError as e:
    print(e)
except Exception as e:
    print(f"An error occurred: {e}")



