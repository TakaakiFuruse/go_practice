package cipher

func rot13(s string) string {

	resSt := ""
	for i := 0; i < len(s); i++ {
		code := s[i]
		if string(s[i]) == "!" || string(s[i]) == " " || string(s[i]) == "." || string(s[i]) == "?" {

			resSt += string(s[i])
		} else if 65 <= code && code <= 77 {
			code += 13
			resSt += string(code)
		} else if 77 < code {
			code = 65 + (12 - (90 - code))
			resSt += string(code)
		}
	}

	return resSt
}
