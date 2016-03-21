package bouncer

import "fmt"

type any []interface{}

func bouncer(ar any) any {
	result := any{}
	for _, value := range ar {
		fmt.Println(value)
		if value != false {
			result = append(result, value)
		}
	}
	return result
}
