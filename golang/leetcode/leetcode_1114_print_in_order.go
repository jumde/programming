package main

import (
	"fmt"
	"sync"
)

type InOrderPrint struct {
	firstDone chan struct {}
	secondDone chan struct {}
}

func NewInOrderPrint() *InOrderPrint {
	return &InOrderPrint{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *InOrderPrint) First(printFirst func()) {
	printFirst()
	close(f.firstDone) // signal that first() is done
}

func (f *InOrderPrint) Second(printSecond func()) {
	<-f.firstDone      // wait for first() to complete
	printSecond()
	close(f.secondDone) // signal that second() is done
}

func (f *InOrderPrint) Third(printThird func()) {
	<-f.secondDone     // wait for second() to complete
	printThird()
}

func main() {
	var wg sync.WaitGroup
	inOrderPrint := NewInOrderPrint()
	wg.Add(3)
	go func() {
		defer wg.Done()
		inOrderPrint.Third(func() { fmt.Println("Third") })
	}()
	go func() {
		defer wg.Done()
		inOrderPrint.Second(func() { fmt.Println("Second") })
	}()
	go func() {
		defer wg.Done();
		inOrderPrint.First(func() { fmt.Println("First") })
	}()
	wg.Wait()
}
