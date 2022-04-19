package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		fmt.Println("waiting for Interrupt signal")
		<-ch
		fmt.Println("Interrupt received")
		ctxCancel()
		fmt.Println("Good Bye!")
		os.Exit(0)
	}()

	num := 10
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 1; i <= num; i++ {
		go cpuBound(i, ctx, &wg)
	}

	wg.Wait()
}

func cpuBound(n int, ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		fmt.Println("waiting for done for ", n)
		for {
			select { // its a buffered or unbuffered??
			case <-ctx.Done():
				wg.Done()
				fmt.Println("Good Bye!", n)
				return
			}
		}
	}()

	fmt.Println("CPU bound:", n)
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		fmt.Fprintf(f, ".")
	}
}
