package main

import (
	"fmt"
	"log"

	"github.com/knei-knurow/frames"
	"github.com/urfave/cli/v2"
)

var turnCommand = cli.Command{
	Name:  "turn",
	Usage: "turn motors using mounted servos. Bear in mind that the rover should be moving while turning",
	Subcommands: []*cli.Command{
		{
			Name:  "left",
			Usage: "turn the wheels left",
			Action: func(c *cli.Context) error {
				log.Println("rover turning left")
				return nil
			},
		},
		{
			Name:  "right",
			Usage: "turn the wheels right",
			Action: func(c *cli.Context) error {
				log.Println("rover turning right")
				return nil
			},
		},
	},
	Action: func(c *cli.Context) error {
		log.Printf("no direction passed (left/right)")
		return nil
	},
}

var goCommand = cli.Command{
	Name:  "go",
	Usage: "move rover forward or backward",
	OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
		log.Println("error:", err)
		return nil
	},
	Before: func(c *cli.Context) error {
		log.Println("before go command")
		return nil
	},
	After: func(c *cli.Context) error {
		log.Println("after go command")
		return nil
	},
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "speed",
			Aliases: []string{"s"},
			Value:   0,
			Usage:   "some speed parameter (currently unused)",
		},
	},
	Subcommands: []*cli.Command{
		{
			Name: "forward",
			Before: func(c *cli.Context) error {
				log.Println("before go forward command")
				return nil
			},
			After: func(c *cli.Context) error {
				log.Println("after go forward command")
				return nil
			},
			Usage: "move the rover forward",
			OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
				log.Println("error:", err)
				return nil
			},
			Action: func(c *cli.Context) error {
				log.Println("rover going forward")

				f := frames.Create([2]byte{'M', 'T'}, []byte{'G', 128})
				n, err := port.Write(f)
				if err != nil {
					return fmt.Errorf("failed to write frame to port: %v", err)
				}

				log.Printf("wrote %d-byte frame to port\n", n)

				return nil
			},
		},
		{
			Name:  "backward",
			Usage: "move the rover backward",
			Action: func(c *cli.Context) error {
				log.Println("rover going backward")
				return nil
			},
		},
	},
	Action: func(c *cli.Context) error {
		log.Printf("no direction passed (forward/backward)")
		return nil
	},
}
