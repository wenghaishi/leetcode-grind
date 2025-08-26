def solution(A, B, C):
    result = []
    count = {'a': A, 'b':B, 'c': C}

    while True:
        added = False
        sorted_count = sorted(count.items(), key=lambda x: -x[1])

        for letter, count in sorted_count:

            if len(result) >= 2 and result[-1] == letter and result[-2] == letter:
                continue

            if count == 0:
                continue

            result.append(letter)
            added = True
            break

        if not added:
            break
    

    return ''.join(result)