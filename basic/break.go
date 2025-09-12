package basic

import (
	"context"
	"fmt"
	"sync"
)

func TestLoopBreak() {
	var wg sync.WaitGroup
	ch := make(chan int)
	cb := context.Background()
	ctx, ctxCanFunc := context.WithCancel(cb)
	defer ctxCanFunc()

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
	L:
		for {
			select {
			case <-c.Done():
				break L
			case v, ok := <-ch:
				if !ok {
					break L
				}
				fmt.Println(v)
			}
		}
		fmt.Println("Ending!")
	}(ctx)

	for i := 0; i < 20; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()
}
