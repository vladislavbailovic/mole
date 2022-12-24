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

	job1 := internal.NewLimitedJob(
		"testdata/**/*.txt",
		internal.NewCommand("ls", []string{"-la"}),
		time.Second*24,
	)
	job2 := internal.NewLimitedJob(
		"testdata/nested",
		internal.NewCommand("ls", []string{"-la"}),
		time.Second*2,
	)
	jobs := []*internal.Job{job1, job2}

	var wg sync.WaitGroup
	for _, j := range jobs {
		wg.Add(1)
		go (func(job *internal.Job) {
			job.Watch(ctx)
			wg.Done()
		})(j)
	}

	wg.Wait()
	fmt.Println("-- all done --")
}
