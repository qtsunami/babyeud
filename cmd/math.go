package cmd

import "github.com/spf13/cobra"

var count int
var start int
var end int

var mathCmd = &cobra.Command{
	Use:   "math",
	Short: "generate some math",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	mathCmd.Flags().IntVarP(&count, "count", "c", 50, "Please input questions count.")
	mathCmd.Flags().IntVarP(&start, "start", "s", 1, "Please input the range start number.")
	mathCmd.Flags().IntVarP(&end, "end", "e", 50, "Please input the range end number.")

}
