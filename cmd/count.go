package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/HardDie/gsplitter/internal"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Split by count of files",
	Run: func(cmd *cobra.Command, args []string) {
		files := internal.ReadFiles()

		count, err := cmd.Flags().GetInt("items")
		if err != nil {
			log.Fatal("error get flag items:", err.Error())
		}
		offset, err := cmd.Flags().GetInt("offset")
		if err != nil {
			log.Fatal("error get flag offset:", err.Error())
		}

		if count <= 0 {
			log.Fatal("flag 'items' should be passed")
		}
		if offset < 0 {
			log.Fatal("flag 'offset' can't be less than 0")
		}

		internal.SplitByCount(files, count, offset)
		log.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	countCmd.Flags().IntP("items", "i", 0, "How many items in single directory")
	countCmd.Flags().IntP("offset", "o", 0, "From which id start to create folders")
}
