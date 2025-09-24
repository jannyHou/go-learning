package context

import (
	"context"
	"fmt"
	"time"
)

func UC1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	str := []string{
		"foo",
		"bar",
		"baz",
	}

	ch := make(chan string, 10)

	for _, s := range str {
		go Print(ctx, s, ch)
	}

	// start printing baz
	// start printing bar
	// start printing foo
	// Timed out!
	// end printing foo
	// end printing baz
	// Channel received foo
	// Channel received baz
	// end printing bar
	// Channel received bar
	// select {
	// case <-ctx.Done():
	// 	fmt.Println("Timed out!")
	// }

	for _ = range str {
		r := <-ch
		fmt.Printf("Channel received %s", r)
		fmt.Println()
	}

	// start printing baz
	// start printing bar
	// start printing foo
	// end printing foo
	// end printing baz
	// Channel received foo
	// Channel received baz
	// end printing bar
	// Channel received bar
	// Timed out!
	// select {
	// case <-ctx.Done():
	// 	fmt.Println("Timed out!")
	// }
}

func Print(ctx context.Context, s string, ch chan<- string) {
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Printf("start printing %s", s)
		fmt.Println()
		time.Sleep(time.Second)
		fmt.Printf("end printing %s", s)
		fmt.Println()
		ch <- s
	}
}
