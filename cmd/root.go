package cmd
import (
"fmt"
"github.com/spf13/cobra"
"os"
)


var rootCmd = &cobra.Command{
	Use:   "iptable-tool",
	Short: "iptable-tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("By o1hy")
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

