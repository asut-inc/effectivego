package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

var m *runtime.MemStats

func tracing() {

	trace.Start(os.Stdout)
	defer trace.Stop()
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	<-ch
}

func main() {
	n, err := os.Stderr.Write(([]byte)("End of program"))
	if err != nil {
		return
	}

	fmt.Println(n)
}
