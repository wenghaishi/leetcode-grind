

# recursive brute force
def solution(n):
    count = 0

    def helper(current):
        if current == n:
            count += 1
        elif current > n:
            return
        else:
            helper(current +1)
            helper(current +2)
    
    helper(0)

    return count
