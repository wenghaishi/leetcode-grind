def solution(numRows):
    solution_array = []

    while numRows > 0:
        if len(solution_array) == 0:
            solution_array.push([1])
        elif len(solution_array) == 1:
            solution_array.push([1,1])
        else:
            for 