package leap

const testVersion = 2

func IsLeapYear(y int) bool {
	if y%4 == 0 {
		if y%100 != 0 {
			return true
		} else if (y/100)%4 == 0 {
			return true
		}
	}
	return false
}
