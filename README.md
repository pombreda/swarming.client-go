Swarming.client (Go)
====================

This directory contains the client code to access the Swarming services.

The fact the service is implemented in Go or in python doesn't matter; the
JSON-REST API is language agnostic.

This code is only tested on Ubuntu.

[![GoDoc](https://godoc.org/github.com/maruel/swarming.client-go?status.svg)](https://godoc.org/github.com/maruel/swarming.client-go)
[![Build Status](https://travis-ci.org/maruel/swarming.client-go.svg?branch=master)](https://travis-ci.org/maruel/swarming.client-go)
[![Coverage Status](https://img.shields.io/coveralls/maruel/swarming.client-go.svg)](https://coveralls.io/r/maruel/swarming.client-go?branch=master)


Installation
------------

    go get -u github.com/maruel/swarming.client-go/swarming
    swarming

And use `swarming help <command>` for further help.
