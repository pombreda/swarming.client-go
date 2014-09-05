// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// swarming - client CLI to access the server.
//
// See README.md for more details.
package main

import (
	"github.com/maruel/subcommands"
	"github.com/maruel/subcommands/subcommandstest"
	"log"
	"os"
)

var application = &subcommands.DefaultApplication{
	Name:  "swarming",
	Title: "Client tool to access a swarming server.",
	// Keep in alphabetical order of their name.
	Commands: []*subcommands.Command{
		cmdTrigger,
		cmdCollect,
		cmdRun,
		subcommands.CmdHelp,
		cmdQuery,
	},
}

type SwarmingApplication interface {
	subcommandstest.Application
}

type swarming struct {
	*subcommands.DefaultApplication
	log *log.Logger
}

// Implementes subcommandstest.Application.
func (d *swarming) GetLog() *log.Logger {
	return d.log
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	d := &swarming{application, log.New(application.GetErr(), "", log.LstdFlags|log.Lmicroseconds)}
	os.Exit(subcommands.Run(d, nil))
}
