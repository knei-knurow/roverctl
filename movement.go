package main

import (
	"fmt"
	"log"

	"github.com/knei-knurow/roverctl/requests"
	"github.com/urfave/cli/v2"
)

var goCommand = cli.Command{
	Name:  "go",
	Usage: "move rover forward or backward",
	OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
		log.Println("error:", err)
		return nil
	},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "speed",
			Aliases: []string{"s"},
			Value:   0,
			Usage:   "some speed parameter",
		},
	},
	Subcommands: []*cli.Command{
		{
			Name:  "forward",
			Usage: "move the rover forward",
			OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
				log.Println("error:", err)
				return nil
			},
			Action: func(c *cli.Context) error {
				body := map[string]interface{}{
					"type":      "go",
					"direction": "forward",
					"speed":     c.Int("speed"),
				}

				err := requests.PostRequest(addr+"/move", body)
				if err != nil {
					return fmt.Errorf("make request: %v", err)
				}

				return nil
			},
		},
		{
			Name:  "backward",
			Usage: "move the rover backward",
			Action: func(c *cli.Context) error {
				body := map[string]interface{}{
					"type":      "go",
					"direction": "backward",
					"speed":     c.Int("speed"),
				}

				err := requests.PostRequest(addr+"/move", body)
				if err != nil {
					return fmt.Errorf("make request: %v", err)
				}

				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop the rover",
			Action: func(c *cli.Context) error {
				body := map[string]interface{}{
					"type":      "go",
					"direction": "stop",
					"speed":     255,
				}

				err := requests.PostRequest(addr+"/move", body)
				if err != nil {
					return fmt.Errorf("make request: %v", err)
				}

				return nil
			},
		},
	},
	Action: func(c *cli.Context) error {
		log.Printf("no direction passed (forward/backward)")
		return nil
	},
}

var turnCommand = cli.Command{
	Name:  "turn",
	Usage: "turn motors using mounted servos. Bear in mind that the rover should be moving while turning",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "degrees",
			Aliases: []string{"d"},
			Value:   0,
			Usage:   "degrees",
		},
	},
	Action: func(c *cli.Context) error {
		body := map[string]interface{}{
			"type":    "turn",
			"degrees": c.Int("degrees"),
		}

		err := requests.PostRequest(addr+"/move", body)
		if err != nil {
			return fmt.Errorf("make request: %v", err)
		}

		return nil
	},
}
