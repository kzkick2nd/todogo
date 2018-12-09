package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const srcFile string = ".todo"

func main() {
	var subCmd, option string
	fmt.Scan(&subCmd, &option)
	fmt.Println(todo(subCmd, option))
}

func todo(subCmd, option string) string {
	currentDir, err := os.Getwd()
	if err == nil {
		p := filepath.Join(currentDir, srcFile)
		f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
		}
		defer f.Close()
	}

	switch subCmd {
	case "add":
		add(f, option)
	case "list":
		list(f)
	case "done":
		done(f, option)
	}
	return "Unknown command:" + subCmd
}

func add(buf, option string) error {}

func list() string {}

func done(option string) error {}
