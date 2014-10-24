// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/maruel/subcommands"
)

type commonFlags struct {
	subcommands.CommandRunBase
	ServerURL string
}

// Init initializes common flags.
func (c *commonFlags) Init() {
	c.Flags.StringVar(&c.ServerURL, "server", os.Getenv("SWARMING_SERVER"), "Server URL; required. Set $SWARMING_SERVER to set a default.")
}

// urlToHTTPS ensures the url is https://.
func urlToHTTPS(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if u.Scheme != "" && u.Scheme != "https" {
		return "", errors.New("Only https:// scheme is accepted. It can be omitted.")
	}
	if !strings.HasPrefix(s, "https://") {
		s = "https://" + s
	}
	if _, err = url.Parse(s); err != nil {
		return "", err
	}
	return s, nil
}

// Parse parses the common flags.
func (c *commonFlags) Parse(d SwarmingApplication) error {
	if c.ServerURL == "" {
		return errors.New("Must provide -server")
	}
	s, err := urlToHTTPS(c.ServerURL)
	if err != nil {
		return err
	}
	c.ServerURL = s
	return nil
}

type doubleVar struct {
	Items map[string]string
}

func (d *doubleVar) String() string {
	out := make([]string, 0, len(d.Items))
	for k, v := range d.Items {
		out = append(out, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(out)
	return strings.Join(out, ", ")
}

func (d *doubleVar) Set(s string) error {
	out := strings.SplitN(s, "=", 1)
	if len(out) != 2 {
		return errors.New("Must use foo=bar format")
	}
	d.Items[out[0]] = out[1]
	return nil
}
