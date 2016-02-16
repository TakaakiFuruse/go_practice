package repeatString

type Repeat struct {
	words string
	num   int
}

func (r Repeat) repeat() string {
	var s string

	for i := 1; i <= r.num; i++ {
		s += r.words
	}
	return s
}
