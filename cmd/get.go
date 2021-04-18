package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd *cobra.Command

func init() {
	getCmd = &cobra.Command{
		Use:   "get <resource>",
		Short: "Get live information on given resource",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resource := args[0]

			fmt.Println(resource)
		},
	}
}
