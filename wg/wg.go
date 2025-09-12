package wg

import (
	"context"
	"fmt"
	"sync"
)

func WgTest() {
	var wg sync.WaitGroup
	cb := context.Background()
	ctx, ctxCancelFn := context.WithCancel(cb)

	chs := make([]chan int, 2)
	for i := 0; i < 2; i++ {
		chs[i] = make(chan int)
	}

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			chs[0] <- i
		}
		close(chs[0])
	}(ctx)

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		for {
			select {
			case <-c.Done():
				return
			case v, ok := <-chs[0]:
				fmt.Printf("2nd go, retrieved ok: %v, value i is %d\n", ok, v)
				if !ok {
					fmt.Println("not ok")
					close(chs[1])
					return
				}
				chs[1] <- v
			}
		}
		// for i := range chs[0] {
		// 	fmt.Printf("2nd go, value i is %d\n", i)
		// 	chs[1] <- i
		// }
	}(ctx)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range chs[1] {
			fmt.Printf("3rd go, value i is %d\n", i)
		}
	}()

	go func() {
		wg.Wait()
		ctxCancelFn()
		fmt.Println("All done")
	}()

	wg.Wait()
	fmt.Println("main")

}

// func StageRunner(c context.Context) {
// 	fmt.Println("stage runner triggered")
// 	for {
// 		select {
// 		case <-c.Done():
// 			fmt.Println("Context cancelled")
// 		default:
// 			fmt.Printf("select default triggered")
// 		}
// 	}
// }
