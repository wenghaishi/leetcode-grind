def solution(n,m):
    divisable = 0
    not_divisable = 0

    for i in range(1, n+1):
        if i % m == 0:
            divisable += i
        else:
            not_divisable += i

    return not_divisable - divisable

print(solution(10, 5))