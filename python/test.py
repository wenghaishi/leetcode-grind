def solution(A, B, C):
    result = []
    counts = {'a': A, 'b': B, 'c': C}

    while True:
        # Sort by remaining count (highest first)
        sorted_counts = sorted(counts.items(), key=lambda x: -x[1])
        added = False

        for letter, count in sorted_counts:
            if count == 0:
                continue

            # Avoid three identical letters in a row
            if len(result) >= 2 and result[-1] == letter and result[-2] == letter:
                continue

            result.append(letter)
            counts[letter] -= 1
            added = True
            break

        if not added:
            break  # No letter can be added without breaking rules

    return ''.join(result)


# Examples
print(solution(6, 1, 1))  # "aabaacaa" or similar
print(solution(1, 3, 1))  # "abbcbb" or similar
print(solution(0, 1, 8))  # "cbcbcbc"
