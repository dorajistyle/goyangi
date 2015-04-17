package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dorajistyle/goyangi/config"
	"github.com/mattes/migrate/migrate"
)

func getAbsPath() string {
	var relPath = "./migrations"
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", absPath)
		return ""
	}
	return absPath
}

func migrateUp() {
	absPath := getAbsPath()
	allErrors, ok := migrate.UpSync("mysql://"+config.MysqlDSL(), absPath)
	if !ok {
		for _, error := range allErrors {
			s := error.Error()
			fmt.Printf("Error! type: %T; value: %q\n", s, s)
		}
		// do sth with allErrors slice
	}

}

func migrateDown() {
	absPath := getAbsPath()

	allErrors, ok := migrate.DownSync("mysql://"+config.MysqlDSL(), absPath)
	if !ok {
		for _, error := range allErrors {
			s := error.Error()
			fmt.Println("Error! type: %T; value: %q\n", s, s)
		}
		// do sth with allErrors slice
	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "down" {
			migrateDown()
			return
		}
	}
	migrateUp()
}
