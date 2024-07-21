package test

import (
	"fmt"
	"testing"
	_"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	//y := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	
	/*for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			y = y + 1
		}
	}*/

	//time.Sleep(6 * time.Second)
	fmt.Println("X Counter = ", x)
	//fmt.Println("Y Counter = ", y)
}