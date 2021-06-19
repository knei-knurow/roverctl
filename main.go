package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// IP of machine running roverd
	ip string

	// port on main computer on which roverd is running
	port string
)

var (
	// addr is ip + ":" + port
	addr string
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("roverctl: ")

	ip = os.Getenv("ROVERCTL_ROVER_IP")
	port = os.Getenv("ROVERCTL_ROVERD_PORT")

	addr = "http://" + ip + ":" + port
	log.Println("addr:", addr)
}

func main() {
	app := &cli.App{
		Name:  "roverctl",
		Usage: "control the Knurów rover command line",
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			log.Println("error:", err)
			return nil
		},
		Commands: []*cli.Command{
			&goCommand,
			&turnCommand,
		},
		CommandNotFound: func(c *cli.Context, command string) {
			log.Printf("invalid command '%s'. See 'roverctl --help'\n", command)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
