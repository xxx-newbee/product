package cmd

import (
	"errors"
	"fmt"

	"github.com/xxx-newbee/product/cmd/api"
	"github.com/xxx-newbee/product/cmd/migrate"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "product",
	Short:        "product",
	Long:         "product-srv",
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least 1 arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tips()
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

func tips() {
	usageStr := `欢迎使用商品服务，请使用 -h 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
