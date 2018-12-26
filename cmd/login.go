package cmd

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/stevenxie/quest-cli/internal/config"
	"github.com/stevenxie/quest-cli/internal/interact"
	ess "github.com/unixpickle/essentials"
)

func registerLoginCmd(app *kingpin.Application) {
	loginCmd = app.Command(
		"login",
		"Save Quest login credentials (obfuscates password).",
	)

	// Register flags.
	loginCmd.Flag("id-only", "Only save Quest ID.").Short('i').
		BoolVar(&loginOpts.IDOnly)
	loginCmd.Flag("clear", "Remove saved login credentials.").Short('c').
		BoolVar(&loginOpts.Clear)
}

var (
	loginCmd  *kingpin.CmdClause
	loginOpts struct {
		IDOnly bool
		Clear  bool
	}
)

func login() error {
	if loginOpts.Clear {
		return clearLogin()
	}

	cfg := new(config.Config)
	if err := interact.PromptMissing(cfg, loginOpts.IDOnly); err != nil {
		return err
	}
	fmt.Println()

	path, err := config.Save(cfg)
	if err != nil {
		return ess.AddCtx("saving file", err)
	}
	fmt.Printf("Credentials saved to '%s'.\n", path)
	return nil
}

func clearLogin() error {
	fmt.Println("Removing config file with saved credentials...")

	removed, err := config.Remove()
	if err != nil {
		return err
	}

	if len(removed) == 0 {
		fmt.Println("No config files were found, nothing to do.")
	} else {
		fmt.Println("Done. The following files were removed:")
		for _, path := range removed {
			fmt.Printf(" - %s\n", path)
		}
	}
	return nil
}
