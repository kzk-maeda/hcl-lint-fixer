package main

import (
	"fmt"
	"os"
)

func getArgs() string {
	args := os.Args
	// fmt.Println(args)
	if len(args) != 2 {
		fmt.Println("[ERROR] Set path to variables.tf as arg.")
		os.Exit(1)
	}
	return args[1]
}

func main() {
	// src := "files/variables.tf"
	src := getArgs()
	// ソースファイルに対するバックアップを作成
	createBackupFile(src)
	// ソースファイルから空行を削除したtmpファイルを作成
	tmpSrc := deleteBlankRow(src)
	// variables.tfの修正
	Run(src, tmpSrc)
	// tmpファイルを削除
	deleteTmpFile(tmpSrc)
}
