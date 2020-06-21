package cmd

import (
	"pepa_pig/importer"
	"strings"

	"github.com/spf13/cobra"
)

func importData() *cobra.Command {

	return &cobra.Command{
		Use:   "import-data",
		Short: "start the http-server",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileName := strings.Join(args, " ")
			importer.InsertRedis(fileName)
		},
	}
}
