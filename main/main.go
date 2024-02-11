package main

import (
	"fmt"
	"strings"
)

func RomanToInt(s string) int {
	var RomanMap = map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	result := 0
	for R := range s {
		if R < len(s)-1 && RomanMap[s[R:R+1]] < RomanMap[s[R+1:R+2]] {
			result -= RomanMap[s[R:R+1]]
		} else {
			result += RomanMap[s[R:R+1]]
		}
	}
	return result
}

func IntegerToRoman(I int) string {
	conversions := []struct {
		key   int
		value string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{8, "VIII"},
		{7, "VII"},
		{6, "VI"},
		{5, "V"},
		{4, "IV"},
		{3, "III"},
		{2, "II"},
		{1, "I"},
	}
	var roman strings.Builder
	for _, conversion := range conversions {
		for I >= conversion.key {
			roman.WriteString(conversion.value)
			I -= conversion.key
		}
	}
	return roman.String()
}

func main() {
	var chose int
	fmt.Println("Введите 1 для выбора Арабского калькулятора или 2 для Римского?")
	fmt.Scan(&chose)
	if chose == 1 {
		var x, y int
		var operator1 string
		fmt.Println("Введите выражение арабскими числами в формате x +|-|*|/ y")
		fmt.Scan(&x, &operator1, &y)
		result1 := x + y
		result2 := x - y
		result3 := x * y
		result4 := x / y
		if x > 10 || y > 10 {
			fmt.Println("Значения не могут быть больше 10")
		} else if x < 0 || y < 0 {
			fmt.Println("Значения не могут быть меньше 0")
		} else if x > 0 && y > 0 && x < 11 && y < 11 && operator1 == "+" {
			fmt.Println(result1)
		} else if x > 0 && y > 0 && x < 11 && y < 11 && operator1 == "-" {
			fmt.Println(result2)
		} else if x > 0 && y > 0 && x < 11 && y < 11 && operator1 == "*" {
			fmt.Println(result3)
		} else if x > 0 && y > 0 && x < 11 && y < 11 && operator1 == "/" {
			fmt.Println(result4)
		}

	} else if chose == 2 {
		var r1, r2 string
		var operator2 string
		fmt.Println("Введите выражение римскими числами в формате r1 +|-|*|/ r2")
		fmt.Scan(&r1, &operator2, &r2)
		rome1 := RomanToInt(r1)
		rome2 := RomanToInt(r2)
		resultRome1 := rome1 + rome2
		resultRome2 := rome1 - rome2
		resultRome3 := rome1 * rome2
		resultRome4 := rome1 / rome2
		int1 := IntegerToRoman(resultRome1)
		int2 := IntegerToRoman(resultRome2)
		int3 := IntegerToRoman(resultRome3)
		int4 := IntegerToRoman(resultRome4)
		if rome1 > 10 || rome2 > 10 {
			fmt.Println("Значения не могут быть больше 10")
		} else if rome1 < 0 || rome2 < 0 {
			fmt.Println("Значения не могут быть меньше 0")
		} else if rome1 > 0 && rome2 > 0 && rome1 < 11 && rome2 < 11 && operator2 == "+" {
			fmt.Println(int1)
		} else if rome1 > 0 && rome2 > 0 && rome1 < 11 && rome2 < 11 && operator2 == "-" {
			fmt.Println(int2)
		} else if rome1 > 0 && rome2 > 0 && rome1 < 11 && rome2 < 11 && operator2 == "*" {
			fmt.Println(int3)
		} else if rome1 > 0 && rome2 > 0 && rome1 < 11 && rome2 < 11 && operator2 == "/" {
			fmt.Println(int4)
		}

	}
}
