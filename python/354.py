
def solution(nums):
    total_area = []
    mapping = {}
    for n in nums:
        area = n[0] * n[1]
        total_area.append(area)
    for index, value in enumerate(nums):
        mapping[value] = total_area[index]
    print(mapping)


solution([[1,1], [2,2], [3,3]])