package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initializationCmd = &cobra.Command{
	Use:     "initialization",
	Aliases: []string{"initialize", "init"},
	Short:   "A brief description of your command",
	Long:    `A longer description of your command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initialization called")
	},
}

func init() {
	rootCmd.AddCommand(initializationCmd)
}
