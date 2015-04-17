package script

import (
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/dorajistyle/goyangi/server"
)

func Commands() []cli.Command {
	return []cli.Command{
		// Init application
		{
			Name:  "init",
			Usage: "Init application(go packages / DB migration / Frontend compiler)",
			Action: func(c *cli.Context) {
				commands := InitApp()
				RunScript(commands)
				println("Init script done.")
			},
		},
		// Generate API document
		{
			Name:    "generateAPI",
			Aliases: []string{"ga"},
			Usage:   "Generate API document using swagger",
			Action: func(c *cli.Context) {
				// go|swagger|asciidoc|markdown|confluence
				format := "asciidoc"
				if len(c.Args()) > 0 {
					format = c.Args()[0]
				}
				commands := GenerateAPI(format)
				RunScript(commands)
				formatRegex, _ := regexp.Compile("^(go|swagger|asciidoc|markdown|confluence)$")
				if formatRegex.MatchString(format) {
					println("API document generated. You can find document at ./document/")
				} else {
					println("Invalid -format specified. Must be one of go|swagger|asciidoc|markdown|confluence.")
				}
			},
		},
		// Run server
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run server",
			Action: func(c *cli.Context) {
				db := ""
				if len(c.Args()) > 0 {
					db = c.Args()[0]
				}
				commands := Server(db)
				RunScript(commands)
				server.Run()
				println("Server is running now.")
			},
		},
		// Run test
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "Run test using Ginkgo",
			Action: func(c *cli.Context) {
				ci := ""
				if len(c.Args()) > 0 {
					ci = c.Args()[0]
				}
				commands := Test(ci)
				RunScript(commands)
				println("Test script done.")
			},
		},
		// Run watcher
		{
			Name:    "watcher",
			Aliases: []string{"w"},
			Usage:   "Run watcher",
			Action: func(c *cli.Context) {
				frontend := "canjs"
				if len(c.Args()) > 0 {
					frontend = c.Args()[0]
				}
				commands := Watcher(frontend)
				RunScript(commands)
				println("Watcher is running now.")
			},
		},
	} // end []cli.Command
} // end Commands()
