package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// Variable is terraform module variable
type Variable struct {
	Name        string         `hcl:",label"`
	Description string         `hcl:"description,optional"`
	Sensitive   bool           `hcl:"sensitive,optional"`
	Type        *hcl.Attribute `hcl:"type,optional"`
	Default     *hcl.Attribute `hcl:"default,optional"`
	Options     hcl.Body       `hcl:",remain"`
}

func judgeType(value string) string {
	/*
		value が "" で囲われていたら：string
		value を数値にキャストできたら：number
		value が [] で囲われていたら：list
		value が true/false であれば：bool
	*/

	if strings.HasPrefix(value, "[") {
		value = strings.TrimRight(strings.TrimLeft(value, "["), "]")    // []を削除
		value = strings.Replace(value, " ", "", -1)                     // 空白を削除
		value = strings.TrimSpace(strings.Replace(value, "\n", "", -1)) // 改行を削除
		splited := strings.Split(value, ",")                            // , でsplit
		samplingValue := splited[0]
		subType := judgeType(samplingValue)
		switch subType {
		case "string":
			return "list_string"
		case "number":
			return "list_number"
		case "bool":
			return "list_bool"
		case "map_string":
			return "list_map_string"
		default:
			return "list"
		}

	} else if strings.HasPrefix(value, "{") {
		return "map_string"
	} else if _, err := strconv.Atoi(value); err == nil {
		return "number"
	} else if value == "true" || value == "false" {
		return "bool"
	} else {
		return "string"
	}
}

func appendTypeToBody(typeValue string, body *hclwrite.Body) error {
	typeTokenMap := getTypesToken()
	// fmt.Println(typeValue, typeTokenMap[typeValue])
	_, ok := typeTokenMap[typeValue]
	if !ok {
		return errors.New("Unexpected Value in Type Label")
	}
	body.AppendUnstructuredTokens(typeTokenMap[typeValue])
	return nil
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
	returnBlock := hclwrite.NewBlock(block.Type(), block.Labels())
	blockName := block.Labels()[0]
	// fmt.Println(variable.Name)
	body := block.Body()

	// descriptionがない場合、追記
	if _, ok := body.Attributes()["description"]; !ok {
		description := parseDescription(blockName)
		returnBlock.Body().SetAttributeValue("description", cty.StringVal(description))
	} else {
		description := body.GetAttribute("description")
		descriptionToken := description.BuildTokens(nil)
		returnBlock.Body().AppendUnstructuredTokens(descriptionToken)
	}

	// typeがない場合、追記
	if _, ok := body.Attributes()["type"]; !ok {
		// defaultがない場合、skip
		if _, ok := body.Attributes()["default"]; !ok {
			msg := fmt.Sprintf("This valiable could not be filled type value : %s", blockName)
			fmt.Println(msg)
			return returnBlock
		}

		defaultValue := string(body.GetAttribute("default").Expr().BuildTokens(nil).Bytes())
		typeValue := judgeType(strings.TrimSpace(defaultValue))
		err := appendTypeToBody(typeValue, returnBlock.Body())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		typeValue := body.GetAttribute("type")
		typeValueToken := typeValue.BuildTokens(nil)
		returnBlock.Body().AppendUnstructuredTokens(typeValueToken)
	}

	// defaultを追記
	if _, ok := body.Attributes()["default"]; ok {
		defaultValue := body.GetAttribute("default")
		defaultValueToken := defaultValue.BuildTokens(nil)
		returnBlock.Body().AppendUnstructuredTokens(defaultValueToken)
	}

	returnBlock.SetLabels(block.Labels())
	returnBlock.SetType(block.Type())

	return returnBlock
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
	for _, block := range body.Blocks() {
		new_block := parseVariable(block)

		body.RemoveBlock(block)
		body.AppendBlock(new_block)
		file.Body().AppendNewline() // blockごとに改行
	}

	updated := file.BuildTokens(nil).Bytes()
	output := hclwrite.Format(updated)
	outputFile := "./files/output.tf"
	// fmt.Fprint(os.Stdout, string(output))
	fo, err := os.Create(outputFile)
	defer fo.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fo.Write(output)
	if err != nil {
		log.Fatal(err)
	}

}
