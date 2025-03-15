# 1769. Minimum Number of Operations to Move All Balls to Each Box
# You have n boxes. You are given a binary string boxes of length n, where boxes[i] is '0' if the ith box is empty, and '1' if it contains one ball.

# In one operation, you can move one ball from a box to an adjacent box. Box i is adjacent to box j if abs(i - j) == 1. Note that after doing so, there may be more than one ball in some boxes.

# Return an array answer of size n, where answer[i] is the minimum number of operations needed to move all the balls to the ith box.

# Each answer[i] is calculated considering the initial state of the boxes.

def minOperations(boxes):
    box_arr = [int(box) for box in boxes]
    result = []

    for outer_index, outer_box in enumerate(box_arr):
        count = 0
        for inner_index, inner_box in enumerate(box_arr):
            if (outer_index != inner_index and inner_box != 0):
                if (inner_index > outer_index):
                    count_to_add = inner_index - outer_index
                    count += count_to_add
                elif (outer_index > inner_index):
                    count_to_add = outer_index - inner_index
                    count += count_to_add
        result.append(count)

    return result

def optimised(boxes):
    arr_length = boxes.length

    for i in range(arr_length):
        



print(minOperations("101"))