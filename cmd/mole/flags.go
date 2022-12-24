package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"mole/internal"
	"strings"
	"time"
)

type Config struct {
	Paths    []string                `json:"paths"`
	Maxdepth int                     `json:"maxdepth"`
	Cmd      string                  `json:"cmd"`
	Interval time.Duration           `json:"interval"`
	Timeout  *time.Duration          `json:"timeout"`
	Target   *internal.CommandTarget `json:"target"`
}

type arrayFlags []string

func (x *arrayFlags) String() string {
	return fmt.Sprintf("%#v", *x)
}

func (x *arrayFlags) Set(value string) error {
	*x = append(*x, strings.Trim(value, `'"`))
	return nil
}

func ParseFlags(args []string) Config {
	var (
		paths    arrayFlags
		command  string
		interval time.Duration
		timeout  time.Duration
		// target   string
	)

	cmd := flag.NewFlagSet("mole", flag.ExitOnError)
	cmd.Var(&paths, "path", "Watch these files")
	cmd.StringVar(&command, "command", "", "Run this command")
	cmd.DurationVar(&interval, "interval", internal.DefaultInterval, "Every so often")
	cmd.DurationVar(&timeout, "timeout", 0, "Until")
	cmd.Parse(args)

	cfg := Config{
		Paths:    paths,
		Cmd:      strings.Trim(command, `'"`),
		Interval: interval,
	}
	if timeout > 0 {
		cfg.Timeout = &timeout
	}

	return cfg
}

func PickleConfig(c Config) []byte {
	if out, err := json.Marshal(c); err != nil {
		fmt.Println(err)
		return out
	} else {
		return out
	}
}

func UnpickleConfig(src []byte) Config {
	var c Config
	if err := json.Unmarshal(src, &c); err != nil {
		fmt.Println(err)
	}
	return c
}
