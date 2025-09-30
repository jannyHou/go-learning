package concurrency

import "fmt"

func readNums(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func sq(inputCh <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for n := range inputCh {
			ch <- n * n
		}
		close(ch)
	}()
	return ch
}

func TestPipeline() {
	nums := []int{1, 2, 3, 4, 5}
	inputCh := readNums(nums)
	outputCh := sq(inputCh)
	for o := range outputCh {
		fmt.Println(o)
	}
}
