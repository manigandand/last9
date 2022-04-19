package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

type Handler struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	hand := &Handler{
		ctx:       ctx,
		ctxCancel: ctxCancel,
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		// fmt.Println("waiting for Interrupt signal")
		<-ch
		// fmt.Println("Interrupt received")
		ctxCancel()
		time.Sleep(time.Second * 1)
		fmt.Println("Good Bye!")
		os.Exit(0)
	}()

	router := chi.NewRouter()
	// http://localhost:8090/cpu?num=10
	router.Get("/cpu", hand.cpuHandler)

	log.Println("Starting server on port:", 8090)
	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) cpuHandler(w http.ResponseWriter, r *http.Request) {
	nums := r.URL.Query().Get("num")
	fmt.Println(nums)
	num, err := strconv.Atoi(nums)
	if err != nil {
		fmt.Fprintf(w, "invalid input %s\n", nums)
		return
	}

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 1; i <= num; i++ {
		go cpuBound(i, h.ctx, &wg)
	}

	wg.Wait()
}

func cpuBound(n int, ctx context.Context, wg *sync.WaitGroup) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				fmt.Println("Good Bye!", n)
				return
			}
		}
	}(ctx)

	fmt.Println("Running CPU bound:", n)
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		fmt.Fprintf(f, ".")
	}
}
