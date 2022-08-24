package Scan

import (
	"fmt"
)

func scan() {
	username := "At"
	age := 11
	fmt.Scanln(&username, &age)
	fmt.Println("账号信息为:", username, age)
	fmt.Printf("用户名是:%q,年龄是：%d\n", username, age)
	fmt.Print("用户名sh:%s,年龄是:%d\n", username, age)
	fmt.Println(&username)
}
