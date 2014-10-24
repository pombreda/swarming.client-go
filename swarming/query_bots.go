// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/maruel/subcommands"
)

var cmdQueryBots = &subcommands.Command{
	UsageLine: "bots <options>",
	ShortDesc: "queries the list of known bots by a Swarming server",
	LongDesc:  "Queries the list of known bots by a Swarming Server. Options can be used to filter the output",
	CommandRun: func() subcommands.CommandRun {
		c := &queryBotsRun{}
		c.Init()
		c.Flags.BoolVar(&c.bare, "bare", false, "Shows only the bot id, no meta data")
		c.Flags.Var(&c.dimensions, "dimension", "Filter bots according to dimensions")
		return c
	},
}

type queryBotsRun struct {
	commonFlags
	bare       bool
	dimensions doubleVar
}

func (c *queryBotsRun) main(a queryApplication) error {
	if err := c.Parse(a); err != nil {
		return err
	}
	fmt.Fprintf(a.GetOut(), "TODO: Implement me!\n")
	return nil
}

func (c *queryBotsRun) Run(a subcommands.Application, args []string) int {
	if len(args) != 0 {
		fmt.Fprintf(a.GetErr(), "%s: Unknown arguments.\n", a.GetName())
		return 1
	}
	HandleCtrlC()
	d := a.(queryApplication)
	if err := c.main(d); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
