package internal

import (
	"context"
	"fmt"
	"time"
)

var DefaultInterval time.Duration = time.Second * 2

type ErrorHandling int

const (
	OnErrorBreak ErrorHandling = iota
	OnErrorReport
	OnErrorSilent
)

func ErrorHandlingFromString(s string) ErrorHandling {
	switch s {
	case "report":
		return OnErrorReport
	case "silent":
		return OnErrorSilent
	}
	return OnErrorBreak
}

func (x ErrorHandling) String() string {
	switch x {
	case OnErrorReport:
		return "report"
	case OnErrorSilent:
		return "silent"
	}
	return "break"
}

type Job struct {
	paths    *Pathlist
	command  *Command
	interval time.Duration
	timeout  *time.Duration
	onError  ErrorHandling
	previous *Filelist
}

func NewJob(command *Command) *Job {
	previous := make(Filelist, 0)
	return &Job{
		command:  command,
		interval: DefaultInterval,
		timeout:  nil,
		previous: &previous,
	}
}

func (x *Job) SetPaths(src []string, globDepth int) {
	x.paths = NewPathlist(src, globDepth)
}

func (x *Job) SetInterval(i time.Duration) {
	if i <= 0 {
		i = DefaultInterval
	}
	x.interval = i
}

func (x *Job) SetTimeout(t time.Duration) {
	if t > 0 {
		x.timeout = &t
	}
}

func (x *Job) SetErrorHandling(e ErrorHandling) {
	x.onError = e
}

func (x *Job) Watch(ctx context.Context) {
	tick := time.NewTicker(x.interval)

	var cancel func()
	if x.timeout != nil {
		ctx, cancel = context.WithTimeout(ctx, *x.timeout)
		defer cancel()
	}

work:
	for {
		x.execute()
		select {
		case <-tick.C:
			// fmt.Println("then work some more")
			continue
		case <-ctx.Done():
			// fmt.Println("context done")
			break work
		}
	}

	// fmt.Println("work cleanup code")
}

func (x *Job) execute() {
	paths := ListFiles(x.paths)

	lst := NewFilelist(paths)
	cmp := CompareFilelists(&lst, x.previous)
	x.previous = &lst
	if !cmp.Any() {
		// fmt.Println("no changes, carry on")
		return
	}

	if err := x.command.ExecuteWith(&cmp); err != nil {
		switch x.onError {
		case OnErrorSilent:
			return
		case OnErrorReport:
			fmt.Println(err)
			return
		}
		panic(err)
	}
}
