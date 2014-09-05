// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"github.com/maruel/subcommands"
	"os"
	"sort"
	"strings"
)

// Common flags.
type CommonFlags struct {
	subcommands.CommandRunBase
	ServerURL string
}

func (c *CommonFlags) Init() {
	c.Flags.StringVar(&c.ServerURL, "server", os.Getenv("SWARMING_SERVER"), "Server URL; required. Set $SWARMING_SERVER to set a default.")
}

func (c *CommonFlags) Parse(d SwarmingApplication) error {
	if c.ServerURL == "" {
		return errors.New("Must provide -server")
	}
	// TODO(maruel): Enforce ServerURL is an URL, prefix https:// automatically.
	return nil
}

type DoubleVar struct {
	Items map[string]string
}

func (d *DoubleVar) String() string {
	out := make([]string, 0, len(d.Items))
	for k, v := range d.Items {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(out)
	return strings.Join(out, ", ")
}

func (d *DoubleVar) Set(s string) error {
	out := strings.SplitN(s, "=", 1)
	if len(out) != 2 {
		return errors.New("Must use foo=bar format")
	}
	d.Items[out[0]] = out[1]
	return nil
}
