package main

import (
	"os"

	"github.com/stevenxie/quest-cli/cmd"
)

func main() {
	// Execute app, with args that exclude the caller name.
	cmd.Exec(os.Args[1:])
}
