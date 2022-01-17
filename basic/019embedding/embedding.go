package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base has num: %v", b.num)
}

type container struct {
	//An embedding looks like a field without a name.
	base
	str string
}

func main() {
	c := container{
		base: base{num: 23},
		str:  "jack",
	}
	fmt.Printf("c={num: %v, str: %v}\n", c.num, c.str)
	fmt.Println("also num:", c.base.num)

	//methods of base become methods of a container
	fmt.Println(c.describe())

	//container now implements the describer interface because it embeds base.
	type describer interface {
		describe() string
	}
	var d describer = c
	fmt.Println(d.describe())
}
