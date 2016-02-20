package where

func where(ar []int, num int) int {

	sAr := ar
	if num <= sAr[0] {
		return 0
	}
	if num <= sAr[len(ar)-1] {
		return len(ar)
	}
	for i := 0; i < len(sAr); i++ {
		if sAr[i] <= num && num <= sAr[i+1] && i < len(sAr)-2 {
			return i + 1
		}
	}
	return 0
}
