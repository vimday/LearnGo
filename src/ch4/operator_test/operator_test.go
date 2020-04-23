package operator_test

import (
	"fmt"
	"testing"
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) 不同维度不能比较
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {

	a := 7
	//if a&Readable{//不支持隐式类型转换
	//	t.Log("1->true")
	//}

	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestString(t *testing.T) {
	text1 := "abcdefg"
	fmt.Println(text1[0])
	for index, char := range text1 {
		fmt.Printf("%d %U %c \n", index, char, char)
	}
}
