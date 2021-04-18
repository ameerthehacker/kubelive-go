package main

import (
	"fmt"
	"github.com/ameerthehacker/kubelive/cmd"
	"github.com/ameerthehacker/kubelive/internal"
	"os"
)

func main() {
	err := cmd.KubeliveCmd.Execute()

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n\n", internal.AppName, err.Error())

		os.Exit(1)
	}
}
