package main

import (
	"os"
	"testing"
)

func TestaddExtension(t *testing.T) {
	if filename := addExtension("test", "txt"); filename != "test.txt" {
		t.Error("no much filename")
	}
}

func TestIsExist(t *testing.T) {
	if false != isExist("isexist") {
		t.Error("doesn't work isexist function")
	}
}

func TestMakeDirectory(t *testing.T) {
	const DIRECTORY = "testDirectory"
	makeDirectory(DIRECTORY)
	defer os.RemoveAll(DIRECTORY)
}

func TestCreateFile(t *testing.T) {
	const FILE = "testFile"
	createFile(FILE)
	defer os.Remove(FILE)
}

func TestExecCommand(t *testing.T) {
	const CMD = "echo"
	const ARGS = "\"this is execCommand() test\""
	execCommand(CMD, ARGS)
}
