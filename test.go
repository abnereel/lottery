package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
}

func (p *Person) Say() {
	fmt.Println(p.Name)
}

func main() {
	var p *Person
	p.Say()
	fmt.Println("123")
	sort.Sort(sort.StringSlice{})
}