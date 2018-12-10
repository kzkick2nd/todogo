package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const srcFile string = ".todo"

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	addMesPtr := addCmd.String("task", "", "TODO text. (Required)")
	doneIDPtr := doneCmd.Int("id", "", "Id to done. (Required)")

	if len(os.Args) < 2 {
		fmt.Println("subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
	case "done":
		doneCmd.Parse(os.Args[2:])
	case "list":
		listCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	var msg string
	currentDir, _ := os.Getwd()
	p := filepath.Join(currentDir, srcFile)
	f, _ := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()

	if addCmd.Parsed() {
		if *addMesPtr == "" {
			addCmd.PrintDefaults()
			os.Exit(1)
		}
		add(f, *addMesPtr)
		msg = list(f)
	}
	if doneCmd.Parsed() {
		if *doneIDPtr == "" {
			doneCmd.PrintDefaults()
			os.Exit(1)
		}
		done(f, f, *doneIDPtr)
		msg = list(f)
	}
	if listCmd.Parsed() {
		msg = list(f)
	}

	fmt.Println(msg)
}

func add(w io.Writer, option string) {
	w.Write([]byte(option + "\n"))
}

func list(r io.Reader) string {
	buf, _ := ioutil.ReadAll(r)
	return string(buf)
}

func done(r io.Reader, w io.Writer, option string) {
	scanner := bufio.NewScanner(r)
	var i int
	id, _ := strconv.Atoi(option)
	for scanner.Scan() {
		i++
		if i == id {
			continue
		}
		w.Write([]byte(scanner.Text() + "\n"))
	}
}
