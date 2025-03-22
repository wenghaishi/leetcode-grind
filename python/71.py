def solution(path):
    path = path + "/"
    path_stack = []
    current = ""

    for char in path:
        if char == "/":
            if current == "..":
                if path_stack:
                    path_stack.pop()
            elif current == ".":
                current = ""
            else:
                if len(current) > 0:
                    path_stack.append(current)
            current = ""
        else:
            current = current + char

    solution = ""
    for path in path_stack:
        solution = solution + "/" + path

    print(solution)

    if solution == "":
        return "/"
    return solution


solution("/home///user//")

