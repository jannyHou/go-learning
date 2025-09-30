package main

import (
	"examples/go-learning/concurrency"
	"fmt"
)

func main() {
	fmt.Println("==========  Go learning projects.  ==========")
	// io.IOTest()
	// wg.WgTest()
	// basic.TestLoopBreak()
	// wg.WgCase2()
	// generic.TestStack()
	concurrency.TestPipeline()
}
