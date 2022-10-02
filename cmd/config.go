/*
Copyright © 2022 Zachary Endrulat zendrulat@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	. "github.com/golangast/groundup/src/cliutility"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate config",
	Long:  `Generate config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")

		CreateConfig()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

}
