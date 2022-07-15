package main

import (
	"flag"
	"fmt"
	"jubilant-potato/internal/extract"
	"jubilant-potato/internal/generator"
	"os"
)

var (
	sourceFilePath string
	targetFilePath string
)

func init() {
	flag.StringVar(&sourceFilePath, "source", "", "source go file path")
	flag.StringVar(&targetFilePath, "target", "", "target rust file path")
}

func main() {
	flag.Parse()
	
	file, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("source file:", sourceFilePath, "open err\n error:", err)
		return
	}
	defer file.Close()

	structArrays, err := extract.ReadStructFromGoFile(file)
	if err != nil {
		fmt.Println("extract struct from go file error:", err)
		return
	}

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		fmt.Println("target file:", targetFilePath)
		fmt.Println("create target file error:", err)
		return
	}
	defer targetFile.Close()

	err = generator.GenerateTargetFile(targetFile, structArrays)
	if err != nil {
		fmt.Println("generate target file error:", err)
		return
	}
}
