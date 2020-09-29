package commands

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/credhub-cli/config"
	"code.cloudfoundry.org/credhub-cli/credhub/auth"
	"code.cloudfoundry.org/credhub-cli/util"
)

func init() {
	CredHub.Token = func() {
		cfg := config.ReadConfig()

		if util.TokenIsPresent(cfg.AccessToken) {
			cfg = refreshConfiguration(cfg)
			config.WriteConfig(cfg)
			fmt.Println("Bearer " + cfg.AccessToken)
		} else if os.Getenv("CREDHUB_CLIENT") != "" && os.Getenv("CREDHUB_SECRET") != "" {
			cfg = refreshConfiguration(cfg)
			fmt.Println("Bearer " + cfg.AccessToken)
		} else {
			fmt.Fprint(os.Stderr, "You are not currently authenticated. Please log in to continue.")
		}
		os.Exit(0)
	}
}

func refreshConfiguration(cfg config.Config) config.Config {
	credhubClient, _ := initializeCredhubClient(cfg)
	authObject := credhubClient.Auth
	oauth := authObject.(*auth.OAuthStrategy)
	err := oauth.Refresh()

	if err != nil {
		return cfg
	}

	cfg.AccessToken = oauth.AccessToken()
	cfg.RefreshToken = oauth.RefreshToken()
	return cfg
}
