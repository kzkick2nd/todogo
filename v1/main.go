package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const srcFile string = ".todo"

func main() {
	var subCmd, option string
	fmt.Scan(&subCmd, &option)
	fmt.Println(todo(subCmd, option))
}

func todo(subCmd, option string) string {
	currentDir, _ := os.Getwd()
	p := filepath.Join(currentDir, srcFile)
	f, _ := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()

	var msg string
	switch subCmd {
	case "list":
		msg = list(f)
	case "add":
		add(f, option)
		msg = list(f)
	case "done":
		done(f, f, option)
		msg = list(f)
	default:
		msg = "Unknown command: " + subCmd
	}
	return msg
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
