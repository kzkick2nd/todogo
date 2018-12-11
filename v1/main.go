package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const srcFile string = ".todo"

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	addMesPtr := addCmd.String("task", "", "TODO text. (Required)")
	doneIDPtr := doneCmd.Int("id", 0, "Id to done. (Required)")

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

	// FIXME
	currentDir, _ := os.Getwd()
	p := filepath.Join(currentDir, srcFile)

	if addCmd.Parsed() {
		if *addMesPtr == "" {
			addCmd.PrintDefaults()
			os.Exit(1)
		}
		// FIXME
		f, _ := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		defer f.Close()

		add(f, *addMesPtr)
	}
	if doneCmd.Parsed() {
		if *doneIDPtr == 0 {
			doneCmd.PrintDefaults()
			os.Exit(1)
		}
		// FIXME
		r, _ := os.OpenFile(p, os.O_RDONLY, 0666)
		defer r.Close()

		t := done(r, *doneIDPtr)
		err := ioutil.WriteFile(p, []byte(t), 0644)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
	}
	// FIXME
	f, _ := os.OpenFile(p, os.O_RDONLY, 0666)
	defer f.Close()

	fmt.Println(list(f))
}

func add(w io.Writer, option string) {
	w.Write([]byte(option + "\n"))
}

func list(r io.Reader) string {
	buf, _ := ioutil.ReadAll(r)
	return string(buf)
}

func done(r io.Reader, id int) string {
	scanner := bufio.NewScanner(r)
	var i int
	var t string
	for scanner.Scan() {
		i++
		if i == id {
			continue
		}
		t = scanner.Text() + "\n"
	}
	return t
}
