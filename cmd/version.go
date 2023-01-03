package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// variables set from main
var (
	Commit  *string
	Date    *string
	Version *string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-cli %s (%s) %s\n", *Version, *Commit, convertStringToTime(Date))
	},
}

func convertStringToTime(date *string) time.Time {
	i, err := strconv.ParseInt(*date, 10, 64)
	if err != nil {
		return time.Time{}
	}
	t := time.Unix(i, 0).UTC()
	return t
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
