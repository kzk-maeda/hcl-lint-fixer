package main

import (
	"bufio"
	"log"
	"os"
)

func deleteBlankRow(path string) string {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	var outputText string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			// fmt.Println(s)
			continue // 空白行をスキップ
		}
		outputText += s + "\n"
	}
	// fmt.Println(outputText)

	outputFileName := path + ".tmp"
	fo, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	outputTextByte := []byte(outputText)
	_, err = fo.Write(outputTextByte)
	if err != nil {
		log.Fatal(err)
	}

	return outputFileName
}

func deleteTmpFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
