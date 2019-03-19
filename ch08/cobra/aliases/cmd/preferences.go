package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var preferencesCmd = &cobra.Command{
	Use:     "preferences",
	Aliases: []string{"prefer", "pref", "prf"},
	Short:   "A brief description of your command",
	Long:    `A longer description of your command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("preferences called")
	},
}

func init() {
	rootCmd.AddCommand(preferencesCmd)
}
