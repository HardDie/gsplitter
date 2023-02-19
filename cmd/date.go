package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/HardDie/gsplitter/internal"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Split files by date",
	Run: func(cmd *cobra.Command, args []string) {
		files := internal.ReadFilesWithDates()

		internal.SplitByDate(files)
		log.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)
}
