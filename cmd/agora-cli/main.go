/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/CarlsonYuan/agora-cli/pkg/cmd/root"
)

func main() {
	if err := mainRun(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func mainRun() error {
	return root.NewCmd().Execute()
}
