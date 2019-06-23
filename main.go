package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/dorajistyle/goyangi/script"
	"github.com/dorajistyle/goyangi/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "Goyangi script tool"
	app.Usage = "run scripts!"
	app.Version = "0.1.0"
	app.Author = "https://github.com/dorajistyle(JoongSeob Vito Kim)"
	app.Commands = script.Commands()
	app.Action = func(c *cli.Context) {
		println("Run Server.")
		server.Run()
	}
	app.Run(os.Args)
}
