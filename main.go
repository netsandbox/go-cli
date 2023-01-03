package main

import "github.com/netsandbox/go-cli/cmd"

var (
	commit  = "abcdefg"
	date    = ""
	version = "v0.0.0"
)

func main() {
	// set variables for the version command
	cmd.Commit = &commit
	cmd.Date = &date
	cmd.Version = &version

	cmd.Execute()
}
