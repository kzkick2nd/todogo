package main

import (
	"bytes"
	"testing"
)

func TestTodo(t *testing.T) {}

func TestAdd(t *testing.T) {
	var buf bytes.Buffer
	input := "testing"

	if err := add(&buf, input); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "add todo: testing", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}

func TestList(t *testing.T) {}

func TestDone(t *testing.T) {}
