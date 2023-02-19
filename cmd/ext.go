package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/HardDie/gsplitter/internal"
)

// extCmd represents the ext command
var extCmd = &cobra.Command{
	Use:   "ext",
	Short: "Split by extension",
	Run: func(cmd *cobra.Command, args []string) {
		files := internal.ReadFiles()

		internal.SplitByExt(files)
		log.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(extCmd)
}
