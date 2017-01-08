package acronym

import "strings"

const testVersion = 2

func Abbreviate(s string) string {
	var st string
	if strings.Contains(s, ":") {
		r := strings.Split(s, ":")
		return r[0]
	}
	r := strings.Split(s, " ")
	for i := 0; i < len(r); i++ {
		if strings.Contains(r[i], "-") {
			t := strings.Split(r[i], "-")
			for _, v := range t {
				p := strings.Split(v, "")
				st += strings.ToUpper(p[0])
			}
		} else {
			var q []string
			q = strings.Split(r[i], "")
			start := "A"
			end := "Z"
			for i, v := range q {
				if i == 0 || (v == strings.ToUpper(v) && v >= start && v <= end) {
					st += strings.ToUpper(v)
				}
			}
		}
	}
	return st
}
