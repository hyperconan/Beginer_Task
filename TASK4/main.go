package main

import (
	"fmt"
)

type tmp struct {
	name string `json:"name"`
}

func (t *tmp) modifyname() {
	(*t).name = "new name"
	fmt.Println(*t)
	(t).name = "new name 1"
	fmt.Println(t)
}

func main() {
	fmt.Println("hello world!")
	//base.Router.Run(":7913")

	t1 := tmp{name: "old name"}
	t1.modifyname()
	fmt.Println(t1.name)
}
