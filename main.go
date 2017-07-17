package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/molsbee/tweet-streamer/command"
	"github.com/molsbee/tweet-streamer/service"
	"github.com/molsbee/tweet-streamer/util"
	"github.com/urfave/cli"
	"os"
)

func main() {
	logger := &util.Logger{logrus.New()}
	credentialsService := service.Credentials{}
	twitterCommands := command.NewTwitterCommands(credentialsService, logger)

	app := cli.NewApp()
	app.Name = "tweet-streamer"
	app.Usage = "command line utility for consuming twitter api"
	app.Before = func(ctx *cli.Context) error {
		credentials, err := credentialsService.Get()
		if err != nil || credentials == nil || !credentials.IsValid() {
			if ctx.Command.Name != command.SET_CREDENTIALS {
				return fmt.Errorf("please setup credentials before using")
			}
		}

		twitterCommands.InitializeTwitterAPI()
		return nil
	}
	app.Commands = twitterCommands.GetCommands()
	app.Run(os.Args)
}
