package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load .env file:", err)
	}

	ip = os.Getenv("IP")
	port = os.Getenv("PORT")

	addr = ip + ":" + port
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
