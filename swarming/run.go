// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/maruel/subcommands"
)

var cmdRun = &subcommands.Command{
	UsageLine: "run <todo>",
	ShortDesc: "runs task on Swarming and wait for completion",
	LongDesc:  "Runs a task on Swarming and collect results. It's similar to running first the command 'trigger' then 'collect'.",
	CommandRun: func() subcommands.CommandRun {
		c := &runRun{}
		c.Init()
		c.Flags.Var(&c.dimensions, "dimension", "Dimensions to filter on")
		return c
	},
}

type runRun struct {
	commonFlags
	dimensions doubleVar
}

func (c *runRun) main(a SwarmingApplication, args []string) error {
	if err := c.Parse(a); err != nil {
		return err
	}
	fmt.Fprintf(a.GetOut(), "TODO: Implement me!\n")
	return nil
}

func (c *runRun) Run(a subcommands.Application, args []string) int {
	HandleCtrlC()
	d := a.(SwarmingApplication)
	if err := c.main(d, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
