package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func printHelp() {
	fmt.Printf("Usage: go-mod-compare [mod file 1] [mod file 2]\n")
}

func parseModFile(fileName string) (*modfile.File, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return modfile.ParseLax(fileName, data, nil)
}

func compareModFile(file1 *modfile.File, file2 *modfile.File) {
	file1Requires := make(map[string]*modfile.Require, len(file1.Require))
	for _, req := range file1.Require {
		file1Requires[req.Mod.Path] = req
	}
	for _, req := range file2.Require {
		if req1, ok := file1Requires[req.Mod.Path]; ok {
			if req1.Mod.Version != req.Mod.Version {
				fmt.Printf("%s: %s > %s\n", req.Mod.Path, req1.Mod.Version, req.Mod.Version)
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		printHelp()
		os.Exit(1)
	}

	fileName1 := args[0]
	fileName2 := args[1]
	// fmt.Printf("Compare two go.mod file: %s and %s\n", fileName1, fileName2)

	file1, err := parseModFile(fileName1)
	if err != nil {
		fmt.Printf("error parsing file %s\n", fileName1)
		os.Exit(1)
	}

	file2, err := parseModFile(fileName2)
	if err != nil {
		fmt.Printf("error parsing file %s\n", fileName2)
		os.Exit(1)
	}

	compareModFile(file1, file2)
}
