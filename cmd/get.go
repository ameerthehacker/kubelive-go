package cmd

import "github.com/spf13/cobra"

var getCmd *cobra.Command

func init() {
	getCmd = &cobra.Command{
		Use:   "get <resource>",
		Short: "Get live information on given resource",
	}
}
