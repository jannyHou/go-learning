package wg

import (
	"fmt"
	"sync"
)

func WgTest() {
	var wg sync.WaitGroup
	ich := make(chan int)

	wg.Add(5)

	go func() {
		fmt.Println("go routine")
		for v := range ich {
			fmt.Println(v)
			wg.Done()
		}
	}()

	for i := 0; i < 5; i++ {
		ich <- i
	}

	wg.Wait()
	fmt.Println("main")
}
