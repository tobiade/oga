package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tobiade/oga/lang"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run oga source code",
	Args:  cobra.ExactArgs(1),
	RunE:  Run,
}

func Run(cmd *cobra.Command, args []string) error {
	p := lang.NewDefaultSourceProvider(args[0])
	lang.RunSourceCode(p)
	return nil
}
