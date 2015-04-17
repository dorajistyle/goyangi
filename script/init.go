package script

import (
	"log"
	"os"
)

func InitApp() []string {
	commands := make([]string, 8)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	commands = append(commands, "SCRIPTPATH="+dir)
	commands = append(commands, "goop install")
	commands = append(commands, "cd \"$SCRIPTPATH/migrate\"")
	commands = append(commands, "goop install")
	commands = append(commands, GoopGoRun+"migrate.go")
	commands = append(commands, "cd \"$SCRIPTPATH/frontend/canjs/compiler\"")
	commands = append(commands, "npm install --save-dev")
	commands = append(commands, "gulp")
	return commands
}
