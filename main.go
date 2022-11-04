package main

import "github.com/kzk-maeda/hcl-lint-fixer/lintfixer"

func main() {
	src := "variables.tf"
	fixer := lintfixer.ValiableFixer(src)
	fixer.Run()
}
