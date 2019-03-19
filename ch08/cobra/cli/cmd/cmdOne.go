package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// cmdOneCmd represents the cmdOne command
var cmdOneCmd = &cobra.Command{
	Use:   "cmdOne",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmdOne called!")
		number, _ := cmd.Flags().GetInt("number")
		fmt.Println("Going to use number", number)
		fmt.Printf("Square: %d\n", number*number)
	},
}

func init() {
	rootCmd.AddCommand(cmdOneCmd)
	cmdOneCmd.Flags().Int("number", 0, "A help for number")
}
