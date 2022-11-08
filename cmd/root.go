package cmd

import (
	"fmt"
	"sbit/cmd/action"

	"github.com/spf13/cobra"
)

// CommandEngine is structure of cli
type Command struct {
	rootCmd *cobra.Command
}

// NewCommandEngine: the command line boot loader
func NewCommand() *Command {
	var rootCMD = &cobra.Command{
		Use:   "sbit",
		Short: "sbit service command line",
		Long:  "sbit service command line",
	}

	return &Command{
		rootCmd: rootCMD,
	}
}

// Run all command line
func (c *Command) Run() {
	var rootCommands = []*cobra.Command{
		{
			Use:   "serve",
			Short: "Run sbit interactor service",
			Long:  "Run sbit interactor service",
			PreRun: func(cmd *cobra.Command, args []string) {
				//load config
				//config.LoadConfig()

				//show display text
				fmt.Println("SBIT")
			},
			Run: func(cmd *cobra.Command, args []string) {
				action.RunRest()
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				//set post run

			},
		},
	}

	for _, command := range rootCommands {
		c.rootCmd.AddCommand(command)
	}

	err := c.rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

// GetRoot the command line service
func (c *Command) GetRoot() *cobra.Command {
	return c.rootCmd
}
