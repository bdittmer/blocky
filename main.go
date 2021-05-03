package main

import (
	"os"

	"github.com/bdittmer/blocky/cmd"
)

func main() {
	cmd.Execute()
	os.Exit(0)
}
