/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package root

import (
	"os"

	"github.com/CarlsonYuan/agora-cli/pkg/cmd/chat"
	cfgCmd "github.com/CarlsonYuan/agora-cli/pkg/cmd/config"
	"github.com/CarlsonYuan/agora-cli/pkg/config"
	"github.com/CarlsonYuan/agora-cli/pkg/version"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var cfgPath = new(string)

func NewCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "agora-cli <command> <subcommand> [flags]",
		Short: "Agora CLI",
		Long:  "Interact with your Agora applications easily",
		Example: heredoc.Doc(`
	
		`),
		Version: version.FmtVersion(),
	}

	fl := root.PersistentFlags()
	fl.String("app", "", "[optional] Application name to use as it's defined in the configuration file")
	fl.StringVar(cfgPath, "config", "", "[optional] Explicit config file path")

	root.AddCommand(
		cfgCmd.NewRootCmd(),
		chat.NewRootCmd(),
	)

	cobra.OnInitialize(config.GetInitConfig(root, cfgPath))

	root.SetOut(os.Stdout)

	return root
}
