package main

func longestConseqSeq(s []int)int {

	if len(s) < 1 {
		return 0
	}

	set := make(map[int]struct{})

	for _, i := range s {
		set[i] = struct{}{}
	}

	best := 1

	for k := range set {
		_, prevExists := set[k-1]
		if prevExists {
			continue
		} 

		current := 1
		j := k
		for {
			_, exists := set[j+1]

			if exists {
				current++
				j++
			} else {
				if current > best {
					best = current
				}

				break
			}
		}

	}

	return best
}
