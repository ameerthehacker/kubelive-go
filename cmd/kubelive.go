package cmd

import "github.com/spf13/cobra"

var KubeliveCmd *cobra.Command

func init() {
	KubeliveCmd = &cobra.Command{
		Use:   "kubelive [command]",
		Short: "Kubectl tool reinvented to be more interactive",
	}

	KubeliveCmd.AddCommand(getCmd)
}
