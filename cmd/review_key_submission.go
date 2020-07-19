package cmd

import (
	"pepa_pig/importer"

	"github.com/spf13/cobra"
)

func ReviewSubmission() *cobra.Command {

	return &cobra.Command{
		Use:   "review",
		Short: "review",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			importer.CheckInReview()
		},
	}
}
