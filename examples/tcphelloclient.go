package main

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
)

type Stub struct {
	Hello func(string) string
}

func main() {
	client := hprose.NewClient("tcp4://127.0.0.1:4321/")
	var stub *Stub
	client.UseService(&stub)
	fmt.Println(stub.Hello("world"))
}
