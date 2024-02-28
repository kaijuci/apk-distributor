package catalog

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apkpublish",
	Short: "apkpublish is a tool to publish apk to a distribution server",
	Long:  `apkpublish is a tool to publish apk to a distribution server. Firebase App Distribution is supported.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
