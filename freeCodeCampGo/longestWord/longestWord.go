package longestWord

import "strings"

func findLongestWord(s string) int {
	ar := strings.Split(s, " ")
	strStore := ar[0]

	for _, value := range ar {
		if len(value) > len(strStore) {
			strStore = value
		}
	}

	return len(strStore)

}
