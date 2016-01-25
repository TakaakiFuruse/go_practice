package mutations

import "strings"

func mutation(ar []string) bool {
	fir := strings.ToLower(ar[0])
	sec := ar[1]
	secAr := strings.Split(sec, "")
	flag := true
	for _, value := range secAr {
		if strings.Index(fir, strings.ToLower(value)) == -1 {
			flag = false
			break
		}
	}
	return flag
}
