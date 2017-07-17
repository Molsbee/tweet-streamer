package command

import (
	"github.com/urfave/cli"
	"net/url"
	"github.com/chimeracoder/anaconda"
	"fmt"
)

func (tc *TwitterCommands) stream() cli.Command {
	return cli.Command{
		Name: "streams",
		Usage: "wrapper for twitter streaming api types [public, user, site]",
		Subcommands: []cli.Command{
			tc.publicStream(),
		},
	}
}

func (tc *TwitterCommands) publicStream() cli.Command {
	return cli.Command{
		Name: "public",
		Usage: "",
		Action: func(ctx *cli.Context) error {
			stream := tc.twitterAPI.PublicStreamFilter(url.Values{
				"track": []string{},
			})
			defer stream.Stop()

			for v := range stream.C {
				tweet, ok := v.(anaconda.Tweet)
				if !ok {
					tc.logger.Warning("received unexpected value of type %T", v)
				}

				fmt.Println(tweet.Text)
			}

			return nil
		},
	}
}
