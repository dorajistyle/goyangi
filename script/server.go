package script

func Server(db string) []string {
	commands := make([]string, 2)
	switch db {
	case "mysql":
		commands = append(commands, "/usr/bin/mysqld_safe &")
		commands = append(commands, "sleep 10s")
	}
	// commands = append(commands, GoopGoRun+"server.go")
	return commands
}
