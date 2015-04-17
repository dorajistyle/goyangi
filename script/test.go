package script

func Test(ci string) []string {
	commands := make([]string, 3)
	command := GoopExec + "ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --trace"
	switch ci {
	case "travis":
		command = command + " --cover --race --compilers=2"
	}
	commands = append(commands, command)
	return commands
}
