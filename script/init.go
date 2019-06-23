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
	commands = append(commands, "dep ensure")
	commands = append(commands, "cd \"$SCRIPTPATH/frontend/vuejs\"")
	commands = append(commands, "yarn")
	return commands
}
