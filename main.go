package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("roverctl: ")
}

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
			Name:  "forward",
			Usage: "move the rover forward",
			Action: func(c *cli.Context) error {
				log.Println("rover going forward")
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

func main() {
	app := &cli.App{
		Name:  "roverctl",
		Usage: "control the Knurów rover command line",
		Action: func(c *cli.Context) error {
			log.Println("no commands passed (roverctl --help to show help)")
			return nil
		},
		Commands: []*cli.Command{
			&goCommand,
			&turnCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
