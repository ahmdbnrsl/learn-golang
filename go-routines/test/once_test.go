package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}
	
	var counter = 0

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			once.Do(func () {
			    counter = 6
			})
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
