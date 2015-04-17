package script

import (
	"log"
	"os"
)

func Watcher(frontend string) []string {
	commands := make([]string, 4)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	commands = append(commands, "SCRIPTPATH="+dir)
	switch frontend {
	case "canjs":
		compilerPath := "frontend/" + frontend + "/compiler"
		commands = append(commands, "cd \"$SCRIPTPATH/"+compilerPath+"\"")
		commands = append(commands, "gulp")
		commands = append(commands, "gulp watch")
	}
	return commands
}
