package wg

import (
	"fmt"
	"sync"
)

func WgTest() {
	var wg sync.WaitGroup

	go func() {
		fmt.Println("go routine")
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
	fmt.Println("main")
}
