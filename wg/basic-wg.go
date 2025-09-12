package wg

import "fmt"

func BasicWgTest() {
	// // It halts deadlock due to no sender
	// ch := make(chan int)
	// value := <-ch
	// fmt.Println(value)
	// // it halts deadlock due to buffered channel full
	// ch2 := make(chan int, 1)
	// ch2 <- 1
	// ch2 <- 2

	// it halts deadlock due to cha and chb waits for each outher as a cycle
	cha := make(chan int)
	chb := make(chan int)

	go func() {
		i := <-cha
		chb <- i
	}()

	go func() {
		j := <-chb
		fmt.Println(j)
	}()

	cha <- 1

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("worker 1 done")
	// }()

	// // worker 2 never calls Done
	// go func() {
	// 	fmt.Println("worker 2 stuck")
	// }()

	// wg.Wait()
	// fmt.Println("all workers finished")
}
