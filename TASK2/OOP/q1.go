package main

import (
	"fmt"
	"math"
)

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	border1 float32
	border2 float32
	border3 float32
}

type Circle struct {
	radius float32
}

func (sha Rectangle) Area() float32 {
	per := sha.Perimeter() / 2
	return float32(math.Sqrt(float64(per * (per - sha.border1) * (per - sha.border2) * (per - sha.border3))))
}

func (sha Rectangle) Perimeter() float32 {
	return sha.border3 + sha.border1 + sha.border2
}

func (sha Circle) Area() float32 {
	return 4 * math.Pi * sha.radius
}

func (sha Circle) Perimeter() float32 {
	return 2 * math.Pi * sha.radius
}

func main() {
	rec := Rectangle{
		border1: 3,
		border2: 4,
		border3: 5,
	}
	fmt.Println("Area:", rec.Area(), " Per:", rec.Perimeter())
	cir := Circle{
		radius: 4,
	}
	fmt.Println("Area:", cir.Area(), " Per:", cir.Perimeter())
}
