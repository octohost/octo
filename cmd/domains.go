// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "Get and set domain names for a specific Docker container.",
	Long:  `Get and set domain names for a specific Docker container.`,
	Run:   startDomains,
}

func startDomains(cmd *cobra.Command, args []string) {
	fmt.Println("octo domains -h")
}

var ()

func init() {
	RootCmd.AddCommand(domainsCmd)
}
