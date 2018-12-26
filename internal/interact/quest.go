package interact

import (
	"github.com/stevenxie/quest-cli/internal/config"
	"github.com/stevenxie/uwquest"
	ess "github.com/unixpickle/essentials"
)

// BuildClient builds and configures a logged-in Quest client.
func BuildClient() (*uwquest.Client, error) {
	c, err := uwquest.NewClient()
	if err != nil {
		return nil, ess.AddCtx("interact: creating Quest client", err)
	}

	cfg, err := config.Load()
	if err != nil {
		return nil, ess.AddCtx("interact: loading config", err)
	}
	if err = PromptMissing(cfg, false); err != nil {
		return nil, err
	}

	Errln("Logging into Quest...")
	err = c.Login(cfg.QuestID, cfg.Password)
	return c, ess.AddCtx("interact: logging into Quest", err)
}
