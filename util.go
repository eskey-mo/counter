package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 拡張子を付与
func addExtension(filename string, extension string) string {
	return filename + "." + extension
}

// ファイルの存在確認
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// ディレクトリ作成
func makeDirectory(dirname string) {
	if err := os.MkdirAll(dirname, 0777); err != nil {
		fmt.Println(dirname + " create failed")
		fmt.Println(err)
		os.Exit(1)
	}
}

// ファイル作成
func createFile(filename string) *os.File {
	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(filename + " create failed")
		os.Exit(1)
	}
	return dst
}

func execCommand(cmd string, args string) {
	exec := exec.Command(cmd, args)
	exec.Stdin = os.Stdin
	exec.Stdout = os.Stdout
	exec.Run()
}
