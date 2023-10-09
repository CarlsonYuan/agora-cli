package chat

import (
	"github.com/CarlsonYuan/agora-cli/pkg/cmd/chat/user"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "Allows you to interact with your Chat applications",
	}
	cmd.AddCommand(user.NewCmds()...)
	return cmd
}
