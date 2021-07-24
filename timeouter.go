package main

import (
	"context"
	"log"
	"time"
)

func main() {
	var (
		workTimeCost  = 6 * time.Second
		cancelTimeout = 5 * time.Second
	)

	ctx, cancel := context.WithCancel(context.Background())

	var (
		data   int
		readCh = make(chan struct{})
	)
	go func() {
		log.Println("blocked to read data")
		// fake long i/o operations
		time.Sleep(workTimeCost)
		data = 10
		log.Println("done read data")

		readCh <- struct{}{}
	}()

	// cancel is called from the other routine
	time.AfterFunc(cancelTimeout, cancel)

	select {
	case <-ctx.Done():
		log.Println("cancelled")
		return
	case <-readCh:
		break
	}

	log.Println("got final data", data)
}
