package main

import (
	"context"
	"mole/internal"
	"os"
	"strings"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := ParseFlags(os.Args[1:])
	var jobs []*internal.Job
	if len(cfg.Paths) == 0 {
		var file string
		if cfg.File != "" {
			file = cfg.File
		} else {
			file = DefaultRcFilename
		}
		jobs = rcfile2JobList(file)
	} else {
		jobs = cfg2JobList(cfg)
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
}

func cfg2JobList(cfg Config) []*internal.Job {
	cmd := strings.Split(cfg.Cmd, " ")
	command := internal.NewCommand(cmd[0], cmd[1:])
	if cfg.Target != nil {
		command.SetTarget(*cfg.Target)
	}

	job := internal.NewJob(command)

	if cfg.Timeout != nil {
		job.SetTimeout(*cfg.Timeout)
	}
	job.SetErrorHandling(cfg.ErrorHandling)
	job.SetInterval(cfg.Interval)
	job.SetPaths(cfg.Paths, cfg.Maxdepth)

	return []*internal.Job{job}
}

func rcfile2JobList(file string) []*internal.Job {
	var jobs []*internal.Job
	cnt, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cfgs := UnpickleConfig(cnt)
	for _, cfg := range cfgs {
		jobs = append(jobs, cfg2JobList(cfg)...)
	}

	return jobs
}
