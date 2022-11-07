package main

func main() {
	src := "files/variables.tf"
	tmpSrc := deleteBlankRow(src)
	Run(tmpSrc)
	deleteTmpFile(tmpSrc)
}
