// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/maruel/interrupt"
	"github.com/maruel/subcommands"
)

var cmdQuery = &subcommands.Command{
	UsageLine: "query <options>",
	ShortDesc: "queries details on Swarming",
	LongDesc:  "Queries various details about a Swarming server.",
	CommandRun: func() subcommands.CommandRun {
		c := &queryRun{}
		c.Init()
		return c
	},
}

type queryRun struct {
	commonFlags
}

func (c *queryRun) main(a SwarmingApplication, args []string) error {
	if err := c.Parse(a); err != nil {
		return err
	}
	fmt.Fprintf(a.GetOut(), "TODO: Implement me!\n")
	return nil
}

// 'query' is itself an application with subcommands.
type queryApplication struct {
	SwarmingApplication
}

func (q queryApplication) GetName() string {
	return q.SwarmingApplication.GetName() + " query"
}

func (q queryApplication) GetCommands() []*subcommands.Command {
	// Keep in alphabetical order of their name.
	return []*subcommands.Command{
		cmdQueryBots,
		subcommands.CmdHelp,
		cmdQueryStartSlave,
	}
}

func (c *queryRun) Run(a subcommands.Application, args []string) int {
	interrupt.HandleCtrlC()
	d := a.(SwarmingApplication)
	// Create an inner application.
	return subcommands.Run(queryApplication{d}, args)
}
