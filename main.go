package main

import (
	"github.com/kzk-maeda/hcl-lint-fixer/lintfixer"
)

func main() {
	src := "variables.tf"
	lintfixer = VariableFixer(src)
	lintfixer.Run()
}
