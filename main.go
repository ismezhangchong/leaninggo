package main

import (
	"fmt"
	Eval "leaninggo/submods"
)

func main() {
	a := "空山新雨后"
	x := 12
	y := 14
	fmt.Println(a)
	fmt.Println("谢谢大佬")
	fmt.Println(Eval.Area())
	fmt.Println(Eval.Timenow())
	fmt.Println("hello")
	z := Eval.Add(x, y)
	fmt.Println(z)
}
