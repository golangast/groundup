package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/zendrulat123/groundup/cmd/watcherut"
)

// watcherCmd represents the watcher command
var watcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "watches the filesys_info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("watcher called")
		Watching()
	},
}

func init() {
	rootCmd.AddCommand(watcherCmd)
}
