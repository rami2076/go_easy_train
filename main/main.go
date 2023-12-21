package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age++ // デリファレンスが不要。(*p).age++ としなくていい。
}

func (p person) nextAge() int {
	return p.age // デリファレンスが不要。(*p).age++ としなくていい。
}

func main() {
	main2()
}

func sub() {
	tom := person{
		name: "Tom",
		age:  15,
	}

	nick := person{
		name: "nick",
		age:  18,
	}

	fmt.Printf("%+v\n", tom)
	fmt.Printf("%+v\n", nick)
	fmt.Printf("%+v\n", tom.nextAge())
	fmt.Printf("%+v\n", nick.nextAge())

	// (&tom).birthday() としなくていい
	// Goが自動的に「変数のアドレス(&)」だと判定する
	tom.birthday()

	fmt.Printf("%+v\n", tom)
	fmt.Printf("%+v\n", nick)
	fmt.Printf("%+v\n", tom.nextAge())
	fmt.Printf("%+v\n", nick.nextAge())
}
