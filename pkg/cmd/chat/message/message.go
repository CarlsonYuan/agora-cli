package message

import (
	"strings"

	agora_chat "github.com/CarlsonYuan/agora-chat-go/v2"
	"github.com/CarlsonYuan/agora-cli/pkg/config"
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewCmds() []*cobra.Command {
	return []*cobra.Command{
		sendCmd(),
	}
}

func sendCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-message --conversation-type [conversation-type] --conversation-id [conversation-id] --text [text] --user [user-id]",
		Short: "Send a message to a conversation",
		Example: heredoc.Doc(`
			# Sends a text message to '123456' conversation of 'users' conversation type
			$ agora-cli chat send-message --conversation-type messaging --conversation-id 123456 --text "Hello World!" --user "user-1"
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := config.GetConfig(cmd).GetClient(cmd)
			if err != nil {
				return err
			}
			cvtType, _ := cmd.Flags().GetString("conversation-type")
			cvtID, _ := cmd.Flags().GetString("conversation-id")
			user, _ := cmd.Flags().GetString("user")
			text, _ := cmd.Flags().GetString("text")

			m := &agora_chat.Message{
				From:        user,
				To:          strings.Split(cvtID, " "),
				MessageType: agora_chat.MessageTypeTxt,
				Body: &agora_chat.MessageBody{
					Msg: text,
				}}

			resp, err := c.Conversation(cvtType, cvtID).SendMessage(cmd.Context(), user, m)
			if err != nil {
				return err
			}

			cmd.Printf("Message successfully sent. Message id: [%v]\n", resp.Data)
			return nil
		},
	}

	fl := cmd.Flags()
	fl.StringP("conversation-type", "t", "", "[required] conversation type such as 'users' or 'chatgroups' or 'chatrooms'")
	fl.StringP("conversation-id", "i", "", "[required] conversation id")
	fl.StringP("user", "u", "", "[required] User id")
	fl.String("text", "", "[required] Text of the message")
	_ = cmd.MarkFlagRequired("conversation-type")
	_ = cmd.MarkFlagRequired("conversation-id")
	_ = cmd.MarkFlagRequired("user")
	_ = cmd.MarkFlagRequired("text")

	return cmd
}
