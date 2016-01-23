package truncate

import "strings"

func truncate(str string, num int) string {
	strAr := strings.Split(str, "")
	if len(str) < num || len(str) == num {
		return str
	}

	if len(str) >= num && num <= 3 {
		sr := strAr[0:num]
		for i := 1; i <= 3; i++ {
			sr = append(sr, ".")
		}
		return strings.Join(sr, "")
	}

	if len(str) >= num {
		sr := strAr[0:num]

		for i := len(sr) - 1; i > len(sr)-4; i-- {
			sr[i] = "."
		}
		return strings.Join(sr, "")
	}
	return ""
}
