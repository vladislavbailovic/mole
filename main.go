package main

import (
	"context"
	"fmt"
	"mole/internal"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job := internal.NewLimitedJob(
		"testdata/**/*.txt",
		[]string{"ls", "-la"},
		4*time.Second,
	)

	var wg sync.WaitGroup
	wg.Add(1)
	go (func(job *internal.Job) {
		job.Watch(ctx)
		wg.Done()
	})(job)

	wg.Wait()
	fmt.Println("-- all done --")
}
func _main() {
	/*
		ctx, cancel := context.WithCancel(context.Background())
		job := NewJob()

		tick := time.NewTicker(24 * time.Second)
		defer cancel()
		go watch(ctx, job)
		<-tick.C
		cancel()

		fmt.Println("all done", job)
	*/
}
