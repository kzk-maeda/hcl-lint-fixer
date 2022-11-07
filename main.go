package main

func main() {
	src := "files/variables.tf"
	// ソースファイルに対するバックアップを作成
	createBackupFile(src)
	// ソースファイルから空行を削除したtmpファイルを作成
	tmpSrc := deleteBlankRow(src)
	// variables.tfの修正
	Run(src, tmpSrc)
	// tmpファイルを削除
	deleteTmpFile(tmpSrc)
}
