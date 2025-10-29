package main

import (
	"fmt"
	"sort"
)
func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)

    var res [][]int

    var helper func(i int, t []int, r int)

    helper = func(i int, t []int, r int) {
        newT := append([]int{}, t...)
        newT = append(newT, candidates[i])

        if candidates[i] == r {
            res = append(res, newT)
        } else {
            for j := i; j < len(candidates); j++ {
                if candidates[j] > r {
                    break
                } else {
                    helper(j, newT, r - candidates[i])
                }
            }
        }
    }

    var start []int

    for k := 0; k < len(candidates); k++ {
        helper(k, start, target)
    }

    return res
}


func main() {
	test1 := []int{2,3,6,7}
	fmt.Println(combinationSum(test1, 7)) 
}