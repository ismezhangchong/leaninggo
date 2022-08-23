package Eval

import (
	"fmt"
)

func Area() int {
	const (
		LENGTH  int = 10
		WIDTH   int = 5
		a, b, c     = 1, false, "str" //多重赋值
	)
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
	return (area)
}
func Timenow() int {
	var time int = 24
	return time
}

// this test
