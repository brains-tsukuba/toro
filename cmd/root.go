package cmd

import (
	"fmt"
	"os"

	"github.com/brains-tsukuba/toro/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toro",
	Short: "Search github trending.",
	Run: func(cmd *cobra.Command, args []string) {
		paths, err := util.Fetch()
		if err != nil {
			_ = fmt.Errorf("Error: %s", err)
			os.Exit(1)
		}
		for i := 0; i < len(paths); i++ {
			fmt.Println(paths[i])
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
