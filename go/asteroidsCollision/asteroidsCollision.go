package main

import "fmt"

func collide(a []int)[]int {
	result := []int{}

	for _, i := range a {

		// if result is empty, append latest
		if len(result) == 0 {
			result = append(result, i)
		} else {
			top := result[len(result) - 1]

			// top and i same direction
			if (top > 0 && i > 0) || (top < 0 && i < 0) {
				result = append(result, i)
			} else if top < 0 && i > 0 {
				// diff direction but no collision
				result = append(result, i)
			} else {
				// diff direction and collision
				for {

					// if current result is empty, append latest i
					if len(result) == 0 {
						result = append(result, i)
						break
					}
					top = result[len(result)-1]
					absI := -i

					if absI > top {
						result = result[:len(result)-1]
					} else if absI < top {
						break
					} else {
						result = result[:len(result)-1]
						break
					}
				}
			}
		}
	}

	return result
}

func main() {
	test1 := []int{3,5,-6,2,-1,4}
	test2 := []int{10,2,-5}

	fmt.Println(collide(test1))
	fmt.Println(collide(test2))

}
