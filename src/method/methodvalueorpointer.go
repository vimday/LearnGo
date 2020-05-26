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

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
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

	sVals := map[int]S{1: {"A"}}

	// 你只能通过值调用 Read
	sVals[1].Read()

	// 这不能编译通过：
	//sVals[1].Write("test")

	ss := S{"B"}
	ss.Write("aa")

	sPtrs := map[int]*S{1: {"A"}}

	// 通过指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")

}

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}
func helloworld() {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	fmt.Println(i, s2Val)
	//  下面代码无法通过编译。因为 s2Val 是一个值，而 S2 的 f 方法中没有使用值接收器
	//   i = s2Vals
}
