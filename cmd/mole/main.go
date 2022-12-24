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
		command := internal.NewCommand(cmd[0], cmd[1:])
		if cfg.Target != nil {
			command.SetTarget(*cfg.Target)
		}

		var job *internal.Job
		if cfg.Timeout != nil {
			job = internal.NewLimitedJob(
				path, command, *cfg.Timeout)
		} else {
			job = internal.NewJob(
				path, command)
		}
		jobs = append(jobs, job)
	}

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
