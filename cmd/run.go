package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tobiade/oga/lang"
)

var runCmd = &cobra.Command{
	Use:          "run",
	Short:        "Run oga source code",
	Args:         cobra.ExactArgs(1),
	RunE:         Run,
	SilenceUsage: true,
}

func init() {
	runCmd.Flags().Bool("pls", false, "force oga to run your code")
}

func Run(cmd *cobra.Command, args []string) error {
	forceRun, err := cmd.Flags().GetBool("pls")
	if err != nil {
		return err
	}
	p := lang.NewDefaultSourceProvider(args[0])
	return lang.RunSourceCode(p, forceRun)
}
