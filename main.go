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
		err = fmt.Errorf("%s: %s", internal.AppName, err.Error())

		_, _ = fmt.Fprint(os.Stderr, err)
	}
}
