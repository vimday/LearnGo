package csp

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Service Done"
}

func otherTask() {
	fmt.Println("otherTask begin...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask end...")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncServiceBuffer() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1) //非阻塞 buffer
	go func() {
		ret := service()
		fmt.Println("Service returned result...")
		retCh <- ret
		fmt.Println("AsyncService exited")
	}()
	return retCh
}

func TestAsyncServiceBuffer(t *testing.T) {
	retCh := AsyncServiceBuffer()
	otherTask()
	fmt.Println(<-retCh)
}

var wg sync.WaitGroup

//wg:=sync.WaitGroup 错误

func AsyncService() chan string {
	wg.Add(1)
	retCh := make(chan string)
	//retCh := make(chan string, 1) //非阻塞 buffer
	go func() {
		defer wg.Done()
		ret := service()
		fmt.Println("Service returned result...")
		retCh <- ret
		fmt.Println("AsyncService exited")
	}()
	//wg.Done()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	wg.Wait()
}
