package goroutine

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("-----------------")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
