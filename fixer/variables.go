package fixer

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// Variable is terraform module variable
type Variable struct {
	Name        string          `hcl:"name,label"`
	Type        hclwrite.Tokens `hcl:"type,attr"`
	Description string          `hcl:"description,attr"`
	Default     hclwrite.Tokens `hcl:"default,attr"`
}

func judgeType(value string) string {
	// value が "" で囲われていたら：string
	// value を数値にキャストできたら：number
	// value が [] で囲われていたら：list
	// value が true/false であれば：bool

	if strings.HasPrefix(value, "[") {
		// fmt.Println("type is list")
		return "list"
	} else if strings.HasPrefix(value, "{") {
		// fmt.Println("type is map")
		return "map"
	} else if _, err := strconv.Atoi(value); err == nil {
		// fmt.Println("type is number")
		return "number"
	} else if value == "true" || value == "false" {
		// fmt.Println("type is bool")
		return "bool"
	} else {
		// fmt.Println("default type")
		return "string"
	}
}

func parseVariable(block *hclwrite.Block) *Variable {
	variable := Variable{
		Name:    block.Labels()[0],
		Default: hclwrite.TokensForValue(cty.StringVal("")),
	}
	fmt.Println(variable.Name)
	body := block.Body()
	// fmt.Println(body.Attributes())

	// descriptionがない場合、追記
	if _, is_exist := body.Attributes()["description"]; !is_exist {
		// fmt.Println("description is not exist")
		body.SetAttributeValue("description", cty.StringVal(variable.Name))
	}

	// typeがない場合、追記
	if _, is_exist := body.Attributes()["type"]; !is_exist {
		// fmt.Println("type is not exist")
		default_value := string(body.GetAttribute("default").Expr().BuildTokens(nil).Bytes())
		type_value := judgeType(strings.TrimSpace(default_value))
		body.SetAttributeValue("type", cty.StringVal(type_value))
	}

	for k, v := range body.Attributes() {
		value := string(v.Expr().BuildTokens(nil).Bytes())
		fmt.Println(k, value)
	}
	return &variable
}

func Run(path string) {
	// path := "variables.tf"
	src, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	file, diags := hclwrite.ParseConfig(src, path, hcl.InitialPos)
	if diags.HasErrors() {
		log.Fatal(err)
	}

	body := file.Body()
	for _, block := range body.Blocks() {
		fmt.Println("---")
		parseVariable(block)

	}
}
