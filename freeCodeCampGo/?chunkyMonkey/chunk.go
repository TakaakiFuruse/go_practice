package chunk

func chunk(ar []int, num int) [][]int {
	resAr := [][]int{}
	tmAr := []int{}

	for i := 0; i < len(ar); i++ {
		if i%num < num-1 {
			tmAr = append(tmAr, ar[i])
		} else {
			tmAr = append(tmAr, ar[i])
			resAr = append(resAr, tmAr)
			tmAr = nil
		}
		if i == len(ar)-1 && tmAr != nil {
			resAr = append(resAr, tmAr)
		}

	}
	return resAr
}
