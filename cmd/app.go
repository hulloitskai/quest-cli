package cmd

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	ess "github.com/unixpickle/essentials"
)

var (
	// Version is the program version. To be generated upon compilation by:
	//     -ldflags "-X github.com/stevenxie/quest-cli/cmd.Version=$(VERSION)"
	//
	// It should match the output of the following command:
	//     git describe --tags | cut -c 2-
	Version = "unset"

	app = kingpin.New(
		"quest",
		"A command line client for the UW Quest Information System.",
	).Version(Version)
)

// Exec runs the root command. It is the application entrypoint.
func Exec(args []string) {
	var err error

	if len(args) == 0 {
		err = interactive()
		goto check
	}

	switch kingpin.MustParse(app.Parse(args)) {
	// Information commands:
	case gradesCmd.FullCommand():
		err = grades()

		// Other commands:
	case loginCmd.FullCommand():
		err = login()
	case interactiveCmd.FullCommand():
		err = interactive()
	}

check:
	if err != nil {
		ess.Die("Error: " + err.Error())
	}
}

// Configure app.
func init() {
	// Customize help, version flag.
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')

	// Register commands.
	registerGradesCmd(app)
	registerLoginCmd(app)
	registerInteractiveCmd(app)
}
