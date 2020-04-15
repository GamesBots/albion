package commands

import (
	"strings"

	"github.com/mvaude/albion/internal/pkg/albion"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "albion",
	Short: "Albion money maker",
	Long:  "Tool for Albion Online",
}

func init() {
	var cmdGet = &cobra.Command{
		Use:   "get",
		Short: "Get information from albion online data",
		Long:  "Get items and prices from albion online data",
	}

	var cmdPrice = &cobra.Command{
		Use:   "price",
		Short: "Get object price from albion online data",
		Long:  "Get items buying and selling orders from albion online data",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, items := range args {
				for _, item := range strings.Split(items, ",") {
					albion.GetPrice(item)
				}
			}
		},
	}

	cmdGet.AddCommand(cmdPrice)
	rootCmd.AddCommand(cmdGet)
}

/*
Execute command line
*/
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}
