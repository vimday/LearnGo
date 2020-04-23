package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstTry(t *testing.T) {
	t.Log(Monday, Tuesday)
	a := 1
	//if a&Readable{//不支持隐式类型转换
	//	t.Log("1->true")
	//}

	t.Log(a&Readable, a&Writable == Writable, a&Executable == Executable)
}
