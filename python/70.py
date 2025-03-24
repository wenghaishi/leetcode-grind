

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

def cached(n):
    cached = {}

    def helper(current):
        if current == n:
            return 1
        if current > n:
            return 0
        if current in cached:
            return cached[current]
        else:
            current_count = helper(current + 1) + helper(current + 2)
            cached[current] = current_count
            return current_count

    solution = helper(0)

    return solution


