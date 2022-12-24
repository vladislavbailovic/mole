package main

import (
	"context"
	"fmt"
	"mole/internal"
	"os"
	"strings"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := ParseFlags(os.Args[1:])
	jobs := make([]*internal.Job, 0, len(cfg.Paths))
	for _, path := range cfg.Paths {
		cmd := strings.Split(cfg.Cmd, " ")
		var job *internal.Job
		if cfg.Timeout != nil {
			job = internal.NewLimitedJob(
				path,
				internal.NewCommand(cmd[0], cmd[1:]),
				*cfg.Timeout)
		} else {
			job = internal.NewJob(
				path,
				internal.NewCommand(cmd[0], cmd[1:]))
		}
		jobs = append(jobs, job)
	}

	// job1 := internal.NewLimitedJob(
	// 	"testdata/**/*.txt",
	// 	internal.NewCommand("ls", []string{"-la"}),
	// 	time.Second*24,
	// )
	// job2 := internal.NewLimitedJob(
	// 	"testdata/nested",
	// 	internal.NewCommand("ls", []string{"-la"}),
	// 	time.Second*2,
	// )
	// jobs := []*internal.Job{job1, job2}

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
