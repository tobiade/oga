package cmd

import "github.com/spf13/cobra"

var abegCmd = &cobra.Command{
	Use:   "abeg",
	Short: "Ask oga nicely and oga might do what you ask.",
}

func init() {
	abegCmd.AddCommand(runCmd)
}
