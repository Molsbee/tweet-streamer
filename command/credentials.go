package command

import (
	"bufio"
	"fmt"
	"github.com/molsbee/tweet-streamer/model"
	"github.com/urfave/cli"
	"os"
	"strings"
)

const SET_CREDENTIALS = "set-credentials"

func (tc *TwitterCommands) setCredentials() cli.Command {
	return cli.Command{
		Name:  SET_CREDENTIALS,
		Usage: "Setup credentials for interacting with Twitter's API",
		Action: func(ctx *cli.Context) error {
			reader := bufio.NewReader(os.Stdin)
			consumerKey := promptUser(reader, "Consumer Key: ")
			consumerSecret := promptUser(reader, "Consumer Secret: ")
			accessToken := promptUser(reader, "Access Token: ")
			accessSecret := promptUser(reader, "Access Secret: ")

			tc.credentials.Save(model.Credentials{
				ConsumerKey:    consumerKey,
				ConsumerSecret: consumerSecret,
				AccessToken:    accessToken,
				AccessSecret:   accessSecret,
			})

			return nil
		},
	}
}

func promptUser(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')

	return strings.Replace(value, "\n", "", -1)
}
