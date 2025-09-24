package basic

import (
	"fmt"
	"time"
)

func TestHeartbeat() {
	done := make(chan bool)
	time.AfterFunc(10*time.Second, func() { close(done) })

	const interval = time.Second
	heartbeat, result := doWork(done, interval)

	for {
		select {
		case v, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println(v)
		case r, ok := <-result:
			if !ok {
				return
			}
			fmt.Println(r)
		}
	}

}

func doWork(done <-chan bool, interval time.Duration) (<-chan bool, <-chan int) {
	heartbeat := make(chan bool)
	result := make(chan int)
	count := 0

	go func() {
		defer close(heartbeat)

		pulse := time.Tick(interval)

		for {
			select {
			case <-done:
				return
			case <-pulse:
				heartbeat <- true
			}
		}
	}()

	go func() {
		for {
			if count == 5 {
				result <- count
				count = 0
			}
			time.Sleep(500 * time.Millisecond)
			count++
		}
	}()

	return heartbeat, result
}
