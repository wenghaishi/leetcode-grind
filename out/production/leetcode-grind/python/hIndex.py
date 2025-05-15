def hIndex(citations):
    citations.sort(reverse=True)  # Sort citations in descending order
    hIndex = 0

    for i, citation in enumerate(citations):
        if citation >= i + 1:
            hIndex = i + 1
        else:
            break

    return hIndex


print(hIndex([1,3,1]))
print(hIndex([3,0,6,1,5]))
