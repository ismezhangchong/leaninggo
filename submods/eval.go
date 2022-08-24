package Eval

import (
	"fmt"
)

func Area() int {
	const (
		LENGTH  int = 12345
		WIDTH   int = 67890
		a, b, c     = 1, false, "str" //多重赋值
	)
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("魅力: %d", area)
	println()
	println(a, b, c)
	return (area)
}
func Timenow() int {
	var time int = 24
	return time
}
func Scan() {
	username := "At"
	age := 11
	fmt.Scanln(&username, &age)
	fmt.Println("账号信息为:", username, age)
	fmt.Printf("用户名是:%q,年龄是：%d\n", username, age)
	fmt.Print("用户名sh:%s,年龄是:%d\n", username, age)
	fmt.Println(&username)
}
