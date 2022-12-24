package main

import (
	"context"
	"fmt"
	"mole/internal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job := internal.NewJob(
		"testdata/**/*.txt",
		[]string{"ls", "-la"},
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
