package main

import (
	"io/ioutil"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func getTypesToken() map[string]hclwrite.Tokens {
	filePath := "./types.tf"
	typeFileSrc, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	typeFile, diags := hclwrite.ParseConfig(typeFileSrc, filePath, hcl.InitialPos)
	if diags.HasErrors() {
		log.Fatal(err)
	}

	typeMap := map[string]hclwrite.Tokens{}
	body := typeFile.Body()
	for _, block := range body.Blocks() {
		labels := block.Labels()
		// fmt.Println(labels)

		typeAttr := block.Body().GetAttribute("type")
		typeAttrToken := typeAttr.BuildTokens(nil)

		typeMap[labels[0]] = typeAttrToken
	}
	// fmt.Println(typeMap)
	return typeMap
}
