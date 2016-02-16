package seek

func destroyer(ar []int, num ...int) []int {
	resAr := []int{}
	numAr := num

	for i := 0; i < len(ar); i++ {
		for _, numValue := range numAr {
			if ar[i] == numValue {
				ar[i] = 0
			}
		}
	}

	for _, value := range ar {
		if value != 0 {
			resAr = append(resAr, value)
		}
	}
	return resAr
}
