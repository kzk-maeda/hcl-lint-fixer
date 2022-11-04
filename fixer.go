package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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
		return "list"
	} else if strings.HasPrefix(value, "{") {
		return "map"
	} else if _, err := strconv.Atoi(value); err == nil {
		return "number"
	} else if value == "true" || value == "false" {
		return "bool"
	} else {
		return "string"
	}
}

func parseDescription(value string) string {
	reg := "[-_]"
	arr := regexp.MustCompile(reg).Split(value, -1)

	// Title Caseに変換
	for l := range arr {
		// 特定文字列の場合はUpper Case
		if arr[l] == "aws" || arr[l] == "vpc" || arr[l] == "cidr" || arr[l] == "az" {
			arr[l] = strings.ToUpper(arr[l])
		} else {
			// その他はTitle Case
			arr[l] = strings.Title(arr[l])
		}
	}
	dest := strings.Join(arr, " ")
	return dest
}

func parseVariable(block *hclwrite.Block) *hclwrite.Block {
	return_block := hclwrite.NewBlock(block.Type(), block.Labels())
	block_name := block.Labels()[0]
	// fmt.Println(variable.Name)
	body := block.Body()

	// descriptionがない場合、追記
	if _, is_exist := body.Attributes()["description"]; !is_exist {
		description := parseDescription(block_name)
		return_block.Body().SetAttributeValue("description", cty.StringVal(description))
	}

	// typeがない場合、追記
	if _, is_exist := body.Attributes()["type"]; !is_exist {
		// defaultがない場合、skip
		if _, is_exist := body.Attributes()["default"]; !is_exist {
			return return_block
		}

		default_value := string(body.GetAttribute("default").Expr().BuildTokens(nil).Bytes())
		type_value := judgeType(strings.TrimSpace(default_value))
		return_block.Body().SetAttributeValue("type", cty.StringVal(type_value))
	}

	return_block.SetLabels(block.Labels())
	return_block.SetType(block.Type())
	// for k, v := range return_block.Body().Attributes() {
	// 	value := string(v.Expr().BuildTokens(nil).Bytes())
	// 	fmt.Println(k, value)
	// }

	return return_block
}

func Run(path string) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	file, diags := hclwrite.ParseConfig(src, path, hcl.InitialPos)
	if diags.HasErrors() {
		log.Fatal(err)
	}

	body := file.Body()
	// var new_body hclwrite.Body
	for _, block := range body.Blocks() {
		new_block := parseVariable(block)

		body.RemoveBlock(block)
		body.AppendBlock(new_block)

	}

	for _, v := range body.Blocks() {
		fmt.Println("---")
		fmt.Println(v.Labels())
		for _, v := range v.Body().Attributes() {
			value := string(v.Expr().BuildTokens(nil).Bytes())
			fmt.Println(value)
		}
	}

}
