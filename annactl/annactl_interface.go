package main

import (
	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlInterfaceCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlInterfaceCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "interface",
		Short: "Interface for Anna's behaviours.",
		Long:  "Interface for Anna's behaviours.",
		Run:   a.ExecAnnactlInterfaceCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlInterfaceTextCmd())

	return newCmd
}

func (a *annactl) ExecAnnactlInterfaceCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlInterfaceCmd")

	cmd.HelpFunc()(cmd, nil)
}
