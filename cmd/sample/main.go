package main

import (
	"fmt"
	"play-godoc/util/outputer"
)

func main() {
	user := "UserA"
	agree := outputer.Hello(user)
	fmt.Printf("content: %+v\n", agree)
}