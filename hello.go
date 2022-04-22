package main

import (
    "fmt"
)

//var message string = "hello world"

var foo, bar, buz string = "foo", "bar", "buz"

func main() {
    //関数の中であればここで宣言できる
    message := "hello world"
    name := "AAA"

    fmt.Printf(message)
    fmt.Printf("さようなら")
    fmt.Printf(foo)
    fmt.Printf(bar)
    fmt.Printf(buz)

    fmt.Println(hello(name))
    fmt.Println(calc(3,3))
}

func hello(name string) string {
	message := fmt.Sprint("Hi, Welcome ", name)
	return message
}

func calc(atai1 int, atai2 int) (add int, sub int) {
    add = atai1 + atai2
    sub = atai1 - atai2
    return
}