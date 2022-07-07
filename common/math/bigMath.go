package math

import (
	"github.com/shopspring/decimal"
	"math"
)

// Add 加
func Add(a, b string) string {
	num1, _ := decimal.NewFromString(a)
	num2, _ := decimal.NewFromString(b)
	d := num1.Add(num2)
	return d.Truncate(8).String()
}

// Sub 减
func Sub(a, b string) string {
	num1, _ := decimal.NewFromString(a)
	num2, _ := decimal.NewFromString(b)
	d := num1.Sub(num2)
	return d.Truncate(8).String()
}

// Mul 乘
func Mul(a, b string) string {
	num1, _ := decimal.NewFromString(a)
	num2, _ := decimal.NewFromString(b)
	d := num1.Mul(num2)
	return d.Truncate(8).String()
}

// Div 除
func Div(a, b string) string {
	num1, _ := decimal.NewFromString(a)
	num2, _ := decimal.NewFromString(b)
	d := num1.Div(num2)
	return d.Truncate(8).String()
}

// UnitTo 值转换 in 原值 i 要转换的位数 为正即缩小数位，为负则增加倍数
func UnitTo(in string, i int) string {
	if in == "" {
		return ""
	}
	fromString, _ := decimal.NewFromString(in)
	result := fromString.Div(decimal.NewFromFloat(math.Pow10(i)))
	return result.Truncate(8).String()
}
