package main

import (
	"bytes"
	"testing"
)

func TestTodo(t *testing.T) {}

func TestAdd(t *testing.T) {
	var buf bytes.Buffer
	option := "testing"
	add(&buf, option)
	if expected, actual := "testing\n", buf.String(); expected != actual {
		t.Errorf("wont %v but got %v", expected, actual)
	}
}

func TestList(t *testing.T) {
	buf := bytes.NewBufferString("testing\n")
	if expected, actual := "testing\n", list(buf); expected != actual {
		t.Errorf("wont %v but got %v", expected, actual)
	}
}

func TestDone(t *testing.T) {
	src := bytes.NewBufferString("1testing\n2testing\n")
	var dest bytes.Buffer
	option := "1"
	done(src, &dest, option)
	if expected, actual := "2testing\n", dest.String(); expected != actual {
		t.Errorf("wont %v but got %v", expected, actual)
	}
}
