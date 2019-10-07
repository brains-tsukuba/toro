package cmd

import (
	"fmt"
	"os"

	"github.com/brains-tsukuba/toro/util"
	"github.com/spf13/cobra"
)

var (
	lang     string
	since    string
	fullpath bool
)

var rootCmd = &cobra.Command{
	Use:   "toro",
	Short: "Search github trending.",
	Run: func(cmd *cobra.Command, args []string) {
		if since != "" && !(since == "daily" || since == "weekly" || since == "monthly") {
			fmt.Println("Error: since option's value is invalid")
			os.Exit(1)
		}
		paths, err := util.Fetch(lang, since)
		if err != nil {
			_ = fmt.Errorf("Error: %s", err)
			os.Exit(1)
		}
		for i := 0; i < len(paths); i++ {
			if fullpath {
				fmt.Println("https://github.com/" + paths[i])
			} else {
				fmt.Println(paths[i])
			}

		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVar(&lang, "lang", "", "filter with programming langage.")
	rootCmd.PersistentFlags().StringVar(&since, "since", "", "filter with date range( \"daily\", \"weekly\", or \"monthly\" )")
	rootCmd.PersistentFlags().BoolVarP(&fullpath, "fullpath", "f", false, "add https://github.com/ to path head")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
