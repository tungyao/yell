package test

import (
	"../../yell"
	"testing"
)

var log = yell.New(yell.Config{
	Path:     "",
	FileName: "err.log",
}, "[LOG]")

func TestYell(t *testing.T) {
	log.Println("test")
}
