package slash

func slasher(ar []int, num int) []int {
	resAr := []int{}
	if len(ar) < num {
		return resAr
	} else if num <= 0 {
		resAr = ar
	} else {
		for ind, el := range ar {
			if num-1 < ind {
				resAr = append(resAr, el)
			}
		}
	}

	return resAr
}
