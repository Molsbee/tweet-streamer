package command

import (
	"github.com/urfave/cli"
	"net/url"
	"github.com/chimeracoder/anaconda"
	"fmt"
	"strings"
)

const (
	STREAMS_COMMAND = "streams"
	PUBLIC_COMMAND = "public"

	track_flag = "track"
	follow_flag = "follow"
	locations_flag = "locations"
)

var streamFlags = []cli.Flag{
	cli.StringFlag{
		Name: track_flag,
		Usage: "comma separated list of strings",
	},
	cli.StringFlag{
		Name: follow_flag,
		Usage: "comma separated list of strings",
	},
	cli.StringFlag{
		Name: locations_flag,
		Usage: "comma separated list of strings",
	},
}

func (tc *TwitterCommands) stream() cli.Command {
	return cli.Command{
		Name: STREAMS_COMMAND,
		Usage: "wrapper for twitter streaming api types [public, user, site]",
		Subcommands: []cli.Command{
			tc.publicStream(),
		},
	}
}

func (tc *TwitterCommands) publicStream() cli.Command {
	return cli.Command{
		Name: PUBLIC_COMMAND,
		Usage: "--track string --follow string --locations string",
		Flags: streamFlags,
		Action: func(ctx *cli.Context) error {
			request, err := convertToRequest(ctx.String(track_flag), ctx.String(follow_flag), ctx.String(locations_flag))
			if err != nil {
				return cli.ShowCommandHelp(ctx, PUBLIC_COMMAND)
			}

			stream := tc.twitterAPI.PublicStreamFilter(request)
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

func convertToRequest(track, follow, locations string) (url.Values, error) {
	if track == "" && follow == "" && locations == "" {
		return nil, fmt.Errorf("no values provided")
	}

	return url.Values{
		"track": strings.Split(track, ","),
		"follow": strings.Split(follow, ","),
		"locations": strings.Split(locations, ","),
	}, nil
}

