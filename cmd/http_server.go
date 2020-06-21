package cmd

import (
	"pepa_pig/server"
	"strings"

	"github.com/spf13/cobra"
)

func httpServerCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "http-server",
		Short: "start the http-server",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			server.StartHTTPServer(strings.Join(args, " "))
		},
	}
}
