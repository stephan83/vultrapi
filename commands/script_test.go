package commands

import (
	. "github.com/stephan83/vultrapi/clients"
	"os"
)

func ExampleScript() {
	c := NewTestClient(200, scripts)
	NewScript().Fexec(os.Stdout, c, []string{"1689"}, "API_KEY")
	// Output:
	// ID		1689
	// NAME		test boot
	// DATE CREATED	2014-12-12 00:38:22 +0000
	// DATE MODIFIED	2014-12-12 00:38:22 +0000
	// SCRIPT		#!/bin/sh
	// echo "test"
}
