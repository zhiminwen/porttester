package main

import (
	"log"
	"os"

	"./cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Printf("error to execute:%v", err)
		os.Exit(1)
	}
}
