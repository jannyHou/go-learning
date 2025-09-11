package wg

import (
	"context"
	"fmt"
	"sync"
)

func WgTest() {
	var wg sync.WaitGroup
	cb := context.Background()
	ctx, cancelFun := context.WithCancel(cb)
	ich := make(chan int)

	wg.Add(1)

	go func() {
		fmt.Println("go routine")
		StageRunner(ctx, ich)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main")
	cancelFun()

}

func StageRunner(c context.Context, ich chan int) {
	fmt.Println("stage runner triggered")
	for {
		select {
		case <-c.Done():
			fmt.Println("Context cancelled")
		case i, ok := <-ich:
			fmt.Printf("ok is %v, value is %v\n", ok, i)
		}
	}
}
