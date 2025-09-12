package wg

import (
	"fmt"
	"sync"
)

func WgCase1() {
	var wg sync.WaitGroup
	ich := make(chan int, 10)

	go func() {
		for i := 0; i < 5; i++ {
			ich <- i
		}

		close(ich)
	}()

	wg.Add(1)

	go func(ch chan int) {
		fmt.Println("go routine")
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}(ich)

	go func() {
		wg.Wait()
		fmt.Println("All done")
	}()

	wg.Wait()
	fmt.Println("main")

}
