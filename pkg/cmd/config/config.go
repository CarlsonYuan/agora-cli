/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"errors"
	"fmt"
	"net/url"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	cfg "github.com/CarlsonYuan/agora-cli/pkg/config"
	"github.com/MakeNowJust/heredoc"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage app configurations",
	}

	cmd.AddCommand(newAppCmd(), removeAppCmd(), listAppsCmd(), setAppDefaultCmd())

	return cmd
}

func newAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Add a new application",
		Long:  "Add a new application which can be used for further operations",
		Example: heredoc.Doc(`
			# Add a new application to the CLI
			$ agora-cli config new
			? What is the name of your app? (eg. prod, staging, testing) testing
			? What is your App ID? 123456#6543321
			? What is your App certificate ? ***********************************
			? Which base URL do you want to use for Chat? https://aXX.chat.agora.io

			Application successfully added. 🚀
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runQuestionnaire(cmd)
		},
	}
}

func removeAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [app-name-1] [app-name-2] [app-name-n]",
		Short: "Remove one or more application.",
		Long:  "Remove one or more application from the configuration file. This operation is irrevocable.",
		Example: heredoc.Doc(`
			# Remove a single application from the CLI
			$ agora-cli config remove staging

			# Remove multiple applications from the CLI
			$ agora-cli config remove staging testing
		`),
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := cfg.GetConfig(cmd)
			for _, appName := range args {
				if err := config.Remove(appName); err != nil {
					return err
				}
				cmd.Printf("[%s] application successfully removed.\n", appName)
			}

			return nil
		},
	}
}

func listAppsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all applications",
		Long:  "List all applications which are configured in the configuration file",
		Example: heredoc.Doc(`
			# List all applications
			$ agora-cli config list
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			t := tabby.NewCustom(w)
			t.AddHeader("", "Name", "App ID", "App Certificate", "BaseURL")

			config := cfg.GetConfig(cmd)

			for _, app := range config.Apps {
				def := ""
				if app.Name == config.Default {
					def = "(default)"
				}
				appCertificate := fmt.Sprintf("**************%v", app.AppCertificate[len(app.AppCertificate)-4:])
				t.AddLine(def, app.Name, app.AppID, appCertificate, app.BaseURL)
			}
			t.Print()
			return nil
		},
	}
}
func setAppDefaultCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "default [app-name]",
		Short: "Set an application as the default",
		Long: heredoc.Doc(`
			Set an application as the default which will be used
			for all further operations unless specified otherwise.
		`),
		Example: heredoc.Doc(`
			# Set an application as the default
			$ agora-cli config default staging
		`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := cfg.GetConfig(cmd)
			return config.SetDefault(args[0])
		},
	}
}

func runQuestionnaire(cmd *cobra.Command) error {
	var newAppConfig cfg.App
	err := survey.Ask(questions(), &newAppConfig)
	if err != nil {
		return err
	}

	config := cfg.GetConfig(cmd)
	err = config.Add(newAppConfig)
	if err != nil {
		return err
	}

	cmd.Println("Application successfully added. 🚀")
	return nil
}

func questions() []*survey.Question {
	return []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is the name of your app? (eg. prod, staging, testing)"},
			Validate: survey.Required,
		},
		{
			Name:     "AppID",
			Prompt:   &survey.Input{Message: "What is your App ID?"},
			Validate: survey.Required,
		},
		{
			Name:     "AppCertificate",
			Prompt:   &survey.Password{Message: "What is your App certificate?"},
			Validate: survey.Required,
		},
		{
			Name: "BaseURL",
			Prompt: &survey.Input{
				Message: "Which base URL do you want to use for Chat?",
			},
			Validate: func(ans interface{}) error {
				u, ok := ans.(string)
				if !ok {
					return errors.New("invalid url")
				}

				_, err := url.ParseRequestURI(u)
				if err != nil {
					return errors.New("invalid url format make sure it matches <scheme>://<host>")
				}
				return nil
			},
		},
	}
}
