package interact

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/stevenxie/quest-cli/internal/config"
	ess "github.com/unixpickle/essentials"
)

var plainTemplates = &promptui.PromptTemplates{
	Valid:   "{{ . | bold }}: ",
	Success: "{{ . | faint }}: ",
}

// PromptMissing prompts the user to fill the missing values in cfg.
func PromptMissing(cfg *config.Config, skipPass bool) error {
	if cfg == nil {
		return errors.New("interact: cannot fill in fields of a nil Config")
	}

	if cfg.QuestID == "" {
		var (
			prompt = promptui.Prompt{
				Label:     "Quest ID",
				Templates: plainTemplates,
			}
			input, err = prompt.Run()
		)
		if err != nil {
			return ess.AddCtx("interact: prompting for Quest ID", err)
		}
		if input == "" {
			return errors.New("interact: Quest ID must not be empty")
		}
		cfg.QuestID = input
	}

	if cfg.Password == "" {
		var (
			prompt = promptui.Prompt{
				Label:     "Quest password",
				Mask:      '*',
				Templates: plainTemplates,
			}
			input, err = prompt.Run()
		)
		if err != nil {
			return ess.AddCtx("interact: prompting for Quest password", err)
		}
		if input == "" {
			return errors.New("interact: password must not be empty")
		}
		cfg.Password = input
	}

	return nil
}
