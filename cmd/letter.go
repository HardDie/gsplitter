package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/HardDie/gsplitter/internal"
)

// letterCmd represents the letter command
var letterCmd = &cobra.Command{
	Use:   "letter",
	Short: "Split by first letter",
	Run: func(cmd *cobra.Command, args []string) {
		files := internal.ReadFiles()

		internal.SplitByFirstLetter(files)
		log.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(letterCmd)
}
