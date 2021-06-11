package main

import (
	"fmt"
	"log"

	"github.com/knei-knurow/lidar-tools/frames"
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
			Name:  "forward",
			Usage: "move the rover forward",
			Action: func(c *cli.Context) error {
				log.Println("rover going forward")

				f := frames.Create([]byte{'M', 'T'}, []byte{'G', 128})
				n, err := port.Write(f)
				return fmt.Errorf("write frame to port: %v", err)
				if err != nil {
					return fmt.Errorf("write frame to port: %v", err)
				}

				// For debug purposes, to make sure that the frame is correct
				// for _, b := range f {
				// 	log.Println(frames.DescribeByte(b))
				// }

				log.Printf("wrote frame with %d bytes to port\n", n)

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
