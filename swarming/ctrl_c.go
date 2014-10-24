// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package main

import (
	"os"
	"os/signal"
	"sync/atomic"
)

// If non-zero, all processing should be interrupted.
var interrupted int32

// Continuously sends true once Ctrl-C was intercepted.
var InterruptedChannel <-chan bool

// The private one to send to.
var interruptedChannel chan<- bool

func init() {
	c := make(chan bool)
	interruptedChannel = c
	InterruptedChannel = c
}

// HandleCtrlC initializes an handler to handle SIGINT, which is normally sent
// on Ctrl-C.
func HandleCtrlC() {
	chanSignal := make(chan os.Signal)

	go func() {
		<-chanSignal
		Interrupt()
	}()

	signal.Notify(chanSignal, os.Interrupt)
}

// Interrupt simulates an interrupt signal. Can be used to stop background
// processing when an error occured and the process should terminates cleanly.
func Interrupt() {
	atomic.StoreInt32(&interrupted, 1)
	go func() {
		for {
			interruptedChannel <- true
		}
	}()
}

// IsInterrupted returns true once an interrupt signal was received.
func IsInterrupted() bool {
	return atomic.LoadInt32(&interrupted) != 0
}
