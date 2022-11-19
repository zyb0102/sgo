package main

import "fmt"

func main() {
	str := "12和对对对AV被女帝"
	fmt.Println("字符串长度", len([]byte(str)))
	fmt.Println("字符串长度2--", len([]rune(str)))

}
