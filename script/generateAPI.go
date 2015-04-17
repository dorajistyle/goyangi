package script

func GenerateAPI(format string) []string {
	commands := make([]string, 2)
	prefix := GoopExec + "swagger -apiPackage=\"github.com/dorajistyle/goyangi/api/v1\" -mainApiFile=\"github.com/dorajistyle/goyangi/api/route.go\" -basePath=\"http://127.0.0.1:3000\" -format=\"" + format + "\" "
	switch format {
	case "go":
		fallthrough
	case "swagger":
		commands = append(commands, prefix+"-output=\"document/\"")
	case "asciidoc":
		commands = append(commands, prefix+"-output=\"document/API.adoc\"")
		commands = append(commands, "asciidoctor -a icons -a toc2 -a stylesheet=github.css -a stylesdir=./stylesheets document/API.adoc")
	case "markdown":
		commands = append(commands, prefix+"-output=\"document/API.md\"")
	case "confluence":
		commands = append(commands, prefix+"-output=\"document/API.confluence\"")
	}
	return commands
}
