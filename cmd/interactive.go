package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/stevenxie/quest-cli/internal/interact"
	ess "github.com/unixpickle/essentials"
)

func registerInteractiveCmd(app *kingpin.Application) {
	interactiveCmd = app.Command(
		"interactive",
		"Run quest-cli in interactive mode.",
	)
}

var interactiveCmd *kingpin.CmdClause

func interactive() error {
	fmt.Println("Running in interactive mode! Type 'help' to see available " +
		"commands.")
	for {
		var command string

		fmt.Print("> ")
		if _, err := fmt.Scanln(&command); err != nil {
			fmt.Println()
			if err == io.EOF {
				return nil
			}
			return ess.AddCtx("reading input", err)
		}
		command = strings.TrimSpace(command)

		var err error
		switch command {
		// Information commands.
		case gradesCmd.FullCommand():
			err = grades()
		case loginCmd.FullCommand():
			err = login()

		// Custom commands.
		case "poll":
			gradesOpts.Poll = true
			err = grades()

			// Other commands.
		case "help":
			interactiveHelp()
		case "quit":
			os.Exit(0)
		default:
			interact.Errf("Unknown command: '%s'\n", command)
		}

		if err != nil {
			return err
		}
	}
}

func interactiveHelp() {
	type HelpEntry struct {
		Name, Help string
	}

	var (
		entries = []HelpEntry{{"help", "Show available commands in " +
			"interactive mode."}}
		models = app.Model().Commands
	)
	for _, model := range models {
		switch model.Name {
		case "help", "interactive":
			continue
		}
		entries = append(entries, HelpEntry{model.Name, model.Help})
	}

	// Add custom entries:
	entries = append(entries, HelpEntry{
		"poll",
		gradesCmd.GetFlag("poll").Model().Help,
	})
	entries = append(entries, HelpEntry{"quit", "Quit interactive mode."})

	fmt.Println("Commands:")
	for _, entry := range entries {
		fmt.Printf("  %s\n    %s\n\n", entry.Name, entry.Help)
	}
}
