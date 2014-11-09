// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/maruel/interrupt"
	"github.com/maruel/subcommands"
)

var cmdTrigger = &subcommands.Command{
	UsageLine: "trigger <todo>",
	ShortDesc: "triggers a task on Swarming",
	LongDesc:  "Triggers a task on Swarming. The task id is then printed, which can then be used to collect results.",
	CommandRun: func() subcommands.CommandRun {
		c := &triggerRun{}
		c.Init()
		c.Flags.Var(&c.dimensions, "dimension", "Dimensions to filter on")
		return c
	},
}

type triggerRun struct {
	commonFlags
	dimensions doubleVar
}

func (c *triggerRun) main(a SwarmingApplication, args []string) error {
	if err := c.Parse(a); err != nil {
		return err
	}
	fmt.Fprintf(a.GetOut(), "TODO: Implement me!\n")
	return nil
}

func (c *triggerRun) Run(a subcommands.Application, args []string) int {
	interrupt.HandleCtrlC()
	d := a.(SwarmingApplication)
	if err := c.main(d, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
