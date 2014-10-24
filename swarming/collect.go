// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/maruel/subcommands"
)

var cmdCollect = &subcommands.Command{
	UsageLine: "collect <task_id>",
	ShortDesc: "collects a task triggered on Swarming",
	LongDesc:  "Collects results from a Swarming task. By default stdout is printed, meta data can be printed via options.",
	CommandRun: func() subcommands.CommandRun {
		c := &collectRun{}
		c.Init()
		c.Flags.BoolVar(&c.metaData, "meta", false, "Shows metadata instead of stdout")
		return c
	},
}

type collectRun struct {
	commonFlags
	metaData   bool
	dimensions doubleVar
}

func (c *collectRun) main(a SwarmingApplication, taskID string) error {
	if err := c.Parse(a); err != nil {
		return err
	}
	fmt.Fprintf(a.GetOut(), "TODO: Implement me!\n")
	return nil
}

func (c *collectRun) Run(a subcommands.Application, args []string) int {
	if len(args) != 1 {
		fmt.Fprintf(a.GetErr(), "%s: Must only provide a task id.\n", a.GetName())
		return 1
	}
	HandleCtrlC()
	d := a.(SwarmingApplication)
	if err := c.main(d, args[0]); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
