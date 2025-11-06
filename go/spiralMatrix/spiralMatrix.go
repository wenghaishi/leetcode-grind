package main

import "fmt"

func spiralMatrix(matrix [][]int) []int {
    var res []int

    for {
        if len(matrix) == 0 {
            break
        }

        // strip top row
        for _, i := range matrix[0] {
            res = append(res, i)
        }

        matrix = matrix[1:]
        
        // strip last column
        if len(matrix) > 0 && len(matrix[0]) > 0 {
            for j := 0; j <len(matrix); j++ {
                length := len(matrix[j])
                res = append(res, matrix[j][length-1])
                matrix[j] = matrix[j][:length -1]
            }
        }

        // strip bottom row
        if len(matrix) > 0 && len(matrix[0]) > 0 {
            lastRow := matrix[len(matrix) -1]

            for k := len(lastRow) -1; k >= 0; k-- {
                res = append(res, matrix[len(matrix) -1][k])
            }

            matrix = matrix[:len(matrix) -1]
        }

        // strip first column
        if len(matrix) > 0 && len(matrix[0]) > 0 {
            for m := len(matrix) -1; m >= 0; m-- {
                res = append(res, matrix[m][0])
                matrix[m] = matrix[m][1:]
            }
        }

    }

    return res
}


func main() {
    // 1. Standard square
    fmt.Println(spiralMatrix([][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    })) // Expected: [1 2 3 6 9 8 7 4 5]

    // 2. Rectangular (more columns)
    fmt.Println(spiralMatrix([][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10,11,12},
    })) // Expected: [1 2 3 4 8 12 11 10 9 5 6 7]

    // 3. Rectangular (more rows)
    fmt.Println(spiralMatrix([][]int{
        {1, 2},
        {3, 4},
        {5, 6},
        {7, 8},
    })) // Expected: [1 2 4 6 8 7 5 3]

    // 4. Single row
    fmt.Println(spiralMatrix([][]int{
        {7, 8, 9, 10},
    })) // Expected: [7 8 9 10]

    // 5. Single column
    fmt.Println(spiralMatrix([][]int{
        {7},
        {9},
        {6},
    })) // Expected: [7 9 6]

    // 6. One element
    fmt.Println(spiralMatrix([][]int{
        {42},
    })) // Expected: [42]

    // 7. Empty matrix
    fmt.Println(spiralMatrix([][]int{})) // Expected: []

    // 8. Matrix where inner rows become empty after pops
    fmt.Println(spiralMatrix([][]int{
        {1, 2, 3},
        {4},
        {5, 6, 7},
    })) // Expected: [1 2 3 6 7 5 4]

    // 9. Tall & skinny
    fmt.Println(spiralMatrix([][]int{
        {1},
        {2},
        {3},
        {4},
        {5},
    })) // Expected: [1 2 3 4 5]

    // 10. Wide & short
    fmt.Println(spiralMatrix([][]int{
        {1,2,3,4,5},
        {6,7,8,9,10},
    })) // Expected: [1 2 3 4 5 10 9 8 7 6]
}