package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scaleval(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v是值类型
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// 值类型 调用接收者也是指针类型的方法
	v.Scale(10)
	// 值类型 调用接收者也是值类型的方法
	fmt.Println(v.Abs())

	// p是指针类型
	p := &Vertex{4, 3}
	// 指针类型 调用接收者是值类型的方法
	fmt.Println(p.Abs())
	// 指针类型 调用接收者也是指针类型的方法
	p.Scale(10)
	fmt.Println(p.Abs())

	pn := &Vertex{4, 3}

	fmt.Println(pn.Abs())

	pn.Scaleval(10)

	fmt.Println(pn.Abs())
}
