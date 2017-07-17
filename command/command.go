package command

import (
	"github.com/chimeracoder/anaconda"
	"github.com/molsbee/tweet-streamer/service"
	"github.com/molsbee/tweet-streamer/util"
	"github.com/urfave/cli"
)

type TwitterCommands struct {
	credentials service.Credentials
	logger      *util.Logger
	twitterAPI  *anaconda.TwitterApi
}

func NewTwitterCommands(credentials service.Credentials, logger *util.Logger) *TwitterCommands {
	return &TwitterCommands{
		credentials: credentials,
		logger:      logger,
	}
}

func (tc *TwitterCommands) InitializeTwitterAPI() error {
	credentials, err := tc.credentials.Get()
	if err != nil {
		return err
	}

	anaconda.SetConsumerKey(credentials.ConsumerKey)
	anaconda.SetConsumerSecret(credentials.ConsumerSecret)
	tc.twitterAPI = anaconda.NewTwitterApi(credentials.AccessToken, credentials.AccessSecret)

	if tc.logger != nil {
		tc.twitterAPI.SetLogger(tc.logger)
	}

	return nil
}

func (tc *TwitterCommands) GetCommands() []cli.Command {
	return []cli.Command{
		tc.setCredentials(),
		tc.stream(),
	}
}
