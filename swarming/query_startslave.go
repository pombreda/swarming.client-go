// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/maruel/interrupt"
	"github.com/maruel/subcommands"
)

var cmdQueryStartSlave = &subcommands.Command{
	UsageLine: "start_slave <options>",
	ShortDesc: "Gets or sets start_slave.py on Swarming",
	LongDesc:  "Gets or sets start_slave.py on Swarming. It must be a valid python file.",
	CommandRun: func() subcommands.CommandRun {
		c := &queryStartSlaveRun{}
		c.Init()
		c.Flags.StringVar(&c.file, "file", "", "Sets a new version of start_slave.py")
		return c
	},
}

type queryStartSlaveRun struct {
	commonFlags
	file string
}

func (c *queryStartSlaveRun) main(a queryApplication) error {
	if err := c.Parse(a); err != nil {
		return err
	}
	fmt.Fprintf(a.GetOut(), "TODO: Implement me!\n")
	return nil
}

func (c *queryStartSlaveRun) Run(a subcommands.Application, args []string) int {
	if len(args) != 0 {
		fmt.Fprintf(a.GetErr(), "%s: Unknown arguments.\n", a.GetName())
		return 1
	}
	interrupt.HandleCtrlC()
	d := a.(queryApplication)
	if err := c.main(d); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
