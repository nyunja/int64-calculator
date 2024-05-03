package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	if !isNum(os.Args[1]) || !isNum(os.Args[3]) {
		return
	}
	value1 := os.Args[1]
	num1 := Atoi(value1)
	operator := os.Args[2]
	if validOperator(operator) == "false" {
		return
	}
	value2 := os.Args[3]
	num2 := Atoi(value2)

	handleCalculation(num1, operator, num2)
}

func validOperator(operator string) string {
	operators := []string{"+", "-", "/", "*", "%"}
	for _, sign := range operators {
		if sign == operator {
			return "true"
		}
	}
	return "false"
}

func handleCalculation(num1 int64, operator string, num2 int64) {
	var result int64
	moduloError := "No modulo by 0"
	divisionError := "No division by 0"
	if num1 >= 9223372036854775807 || num2 >= 9223372036854775807 {
		return
	}
	if num1 <= -9223372036854775807 || num2 <= -9223372036854775807 {
		return
	}
	if num1 == 1 && operator == "/" && num2 == 0 {
		os.Stdout.WriteString(divisionError + "\n")
		return
	}
	if num1 == 1 && operator == "%" && num2 == 0 {
		os.Stdout.WriteString(moduloError + "\n")
		return
	}
	if operator == "/" {
		result = num1 / num2
		str := Itoa(result)
		os.Stdout.WriteString(str + "\n")
		return
	} else if operator == "*" {
		result = num1 * num2
		print(result)
		str := Itoa(result)
		os.Stdout.WriteString(str + "\n")
		return
	} else if operator == "%" {
		result = num1 % num2
		str := Itoa(result)
		os.Stdout.WriteString(str + "\n")
		return
	} else if operator == "+" {
		result = num1 + num2
		str := Itoa(result)
		os.Stdout.WriteString(str + "\n")
		return
	} else if operator == "-" {
		result = num1 - num2
		str := Itoa(result)
		os.Stdout.WriteString(str + "\n")
		return
	}
}

func isNum(s string) bool {
	countSigns := 0
	status := false
	for _, n := range s {
		if n == '-' {
			countSigns++
			if countSigns > 1 {
				status = false
				countSigns = 0
			} else {
				status = true
			}
		} else if n == '+' {
			countSigns++
			if countSigns > 1 {
				status = false
				countSigns = 0
			} else {
				status = true
			}
		} else if n >= '0' && n <= '9' {
			status = true
		} else {
			status = false
		}
	}
	return status
}

func Atoi(s string) int64 {
	var result int64 = 0
	var sign int64 = 1
	for _, n := range s {
		if n == '-' {
			sign = -1
		} else {
			result = result*10 + int64(n-'0')
		}
	}
	return result * sign
}

func Itoa(n int64) string {
	var result string
	var sign string
	if n < 0 {
		sign = "-"
		n = n * -1
	}
	if n >= 0 {
		var rem string
		for i := 0; n/10 > 0; i++ {
			rem = string(n%10 + '0')
			result = rem + result
			n = n / 10
		}
		rem = string(n + '0')
		result = rem + result
	}
	return sign + result
}
