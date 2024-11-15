package main

import (
	"fmt"
	"os"

	"github.com/Terminal/cmd"
)

func main() {
	if err := cmd.RootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}