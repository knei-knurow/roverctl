package main

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tarm/serial"
	"github.com/urfave/cli/v2"
)

var port io.ReadWriteCloser

func init() {
	log.SetFlags(0)
	log.SetPrefix("roverctl: ")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load .env file:", err)
	}

	portName := os.Getenv("PORT_NAME")
	baudRateStr := os.Getenv("BAUD_RATE")

	baudRate, err := strconv.Atoi(baudRateStr)
	if err != nil {
		log.Fatalf("cannot read baud rate: %v\n", err)
	}

	config := &serial.Config{
		Name: portName,
		Baud: baudRate,
	}

	port, err = serial.OpenPort(config)
	if err != nil {
		log.Fatalf("cannot open port %s: %v\n", portName, err)
	}
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
