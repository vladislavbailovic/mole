package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"mole/internal"
	"strings"
	"time"
)

var DefaultRcFilename = "molerc.json"

type Config struct {
	Paths    []string `json:"paths"`
	File     string   `json:"-"`
	Maxdepth int      `json:"maxdepth"`
	Cmd      string   `json:"cmd"`

	RawInterval string        `json:"interval"`
	Interval    time.Duration `json:"-"`

	RawTimeout time.Duration  `json:"timeout"`
	Timeout    *time.Duration `json:"-"`

	RawTarget string                  `json:"target"`
	Target    *internal.CommandTarget `json:"-"`

	RawErrorHandling string                 `json:"error"`
	ErrorHandling    internal.ErrorHandling `json:"-"`
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
		paths       arrayFlags
		command     string
		file        string
		interval    time.Duration
		timeout     time.Duration
		target      string
		errhandling string
	)

	cmd := flag.NewFlagSet("mole", flag.ExitOnError)
	cmd.Var(&paths, "path", "Watch these files")
	cmd.StringVar(&command, "command", "", "Run this command")
	cmd.DurationVar(&interval, "interval", internal.DefaultInterval, "Every so often")
	cmd.DurationVar(&timeout, "timeout", 0, "Until")
	cmd.StringVar(&target, "target", "", "With these args")
	cmd.StringVar(&file, "file", "", "Load config(s) from")
	cmd.StringVar(&errhandling, "errhandling", "", "On error...")
	cmd.Parse(args)

	cfg := Config{
		Paths:            paths,
		Cmd:              strings.Trim(command, `'"`),
		Interval:         interval,
		File:             file,
		RawTimeout:       timeout,
		RawTarget:        target,
		RawErrorHandling: errhandling,
	}

	return hydrateConfig(cfg)
}

func hydrateConfig(cfg Config) Config {
	if cfg.RawTimeout > 0 {
		cfg.Timeout = &cfg.RawTimeout
	}
	if cfg.Interval == 0 {
		cfg.Interval = internal.DefaultInterval
	}
	if cfg.RawTarget != "" {
		tgt := internal.TargetFromString(cfg.RawTarget)
		cfg.Target = &tgt
	}
	if cfg.RawErrorHandling != "" {
		cfg.ErrorHandling = internal.ErrorHandlingFromString(
			cfg.RawErrorHandling)
	}
	if cfg.RawInterval != "" {
		i, err := time.ParseDuration(cfg.RawInterval)
		if err != nil {
			cfg.Interval = internal.DefaultInterval
		} else {
			cfg.Interval = i
		}
	}
	return cfg
}

func PickleConfig(c []Config) []byte {
	out, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return out
}

func UnpickleConfig(src []byte) []Config {
	var cfgs []Config
	if err := json.Unmarshal(src, &cfgs); err != nil {
		panic(err)
	}
	for i, c := range cfgs {
		cfgs[i] = hydrateConfig(c)
	}
	return cfgs
}
