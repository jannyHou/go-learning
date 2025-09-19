package wg

import (
	"context"
	"fmt"
	"sync"
)

func WgCase2() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	cb := context.Background()
	ctx, cancelFunc := context.WithCancel(cb)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	wg.Add(1)
	go func(c context.Context) {
		chi := make(chan int)
		var wgi sync.WaitGroup
		defer wg.Done()

		for j := 0; j < 3; j++ {
			wgi.Add(1)
			go func() {
				defer wgi.Done()
				worker(c, chi, ch2)
			}()
		}

	L:
		for {
			select {
			case <-ctx.Done():
				break L
			case v, ok := <-ch1:
				if !ok {
					close(chi)
					break L
				}
				fmt.Printf("2nd go, v is %v\n", v)
				chi <- v
			}
		}

		wgi.Wait()
		close(ch2)

	}(ctx)

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch2:
				if !ok {
					return
				}
				fmt.Printf("3rd go, v is %v\n", v)
			}
		}
	}(ctx)

	wg.Wait()
	fmt.Println("Ending!")
	cancelFunc()
}

func worker(c context.Context, chin chan int, chout chan int) {
	for {
		select {
		case <-c.Done():
			return
		case v, ok := <-chin:
			if !ok {
				return
			}
			fmt.Printf("internal go, v is %v\n", v)
			chout <- v
		}
	}
}
