package main

func main() {
	src := "files/variables.tf"
	createBackupFile(src)
	tmpSrc := deleteBlankRow(src)
	Run(src, tmpSrc)
	deleteTmpFile(tmpSrc)
}
