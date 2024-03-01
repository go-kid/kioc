package cmd

import (
	"github.com/go-kid/kioc/cmd/scan"
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "kioc",
	Short: "kioc is used to assist in using the go-kid/ioc",
}

func init() {
	Root.AddCommand(scan.Scan)
}
