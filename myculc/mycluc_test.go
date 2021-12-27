package myculc

func plus(num1, num2 int) int {
	return num1 + num2
}

func substract(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}
