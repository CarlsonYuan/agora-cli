package user

import (
	"github.com/CarlsonYuan/agora-cli/pkg/config"
	"github.com/CarlsonYuan/agora-cli/pkg/utils"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewCmds() []*cobra.Command {
	return []*cobra.Command{
		queryCmd(),
	}
}

func queryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-user --id [userID] --output-format [json|tree]",
		Short: "Query user",
		Long: heredoc.Doc(`
			This command allows you to search for user. The 'id' flag is a string,
			and you can check the valid combinations in the official documentation.

			https://docs.agora.io/en/agora-chat/restful-api/user-system-registration?platform=web#querying-a-user
		`),
		Example: heredoc.Doc(`
			# Query for 'user-1'. The results are shown as json.
			$ agora-cli chat query-user --id 'user-d'
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := config.GetConfig(cmd).GetClient(cmd)
			if err != nil {
				return err
			}
			uID, _ := cmd.Flags().GetString("id")
			resp, err := c.QueryUser(cmd.Context(), uID)
			if err != nil {
				return err
			}

			return utils.PrintObject(cmd, resp)
		},
	}

	fl := cmd.Flags()
	fl.StringP("id", "i", "", "[required] query the spicfied user")
	fl.StringP("output-format", "o", "json", "[optional] Output format. Can be json or tree")
	_ = cmd.MarkFlagRequired("filter")

	return cmd
}
