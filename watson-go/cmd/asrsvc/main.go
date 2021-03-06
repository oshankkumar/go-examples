package main

import (
	"fmt"
	"github.com/oshankfriends/go-examples/watson-go/pkg/cmd/asrctl"
	"os"
)

func main() {
	command := asrctl.NewAsrCommand()
	if err := command.Execute(); err != nil {
		fmt.Errorf("%s\n", err)
		os.Exit(1)
	}
}

