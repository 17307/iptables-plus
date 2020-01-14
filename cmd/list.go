package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(resetCmd)
	resetCmd.AddCommand(resetRuleCmd)
	resetCmd.AddCommand(resetCountCmd)
	ruleChain = map[string][]string{
		"OUTPUT":      {"raw", "mangle", "nat", "filter"},
		"PREROUTING":  {"raw", "mangle", "nat"},
		"INPUT":       {"mangle", "nat", "filter"},
		"FORWARD":     {"mangle", "filter"},
		"POSTROUTING": {"mangle", "nat"},
	}
}

var ruleChain map[string][]string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all rules by chain",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		doCmd(args, "-nvL")
	},
}

func doCmd(args []string, f string) {
	// 如果args为空
	if args[0] == "all" {
		// 如果第一个参数为all，则列出所有内容
		for chainName, _ := range ruleChain {
			fmt.Println("\033[31;1m", chainName, "in", ruleChain[chainName], "with", f, "\033[0m")
			for _, j := range ruleChain[chainName] {
				goExecCmd(j, f, chainName)
			}
			if f=="-nvL"{
				fmt.Println()
			}
		}
	} else {
		// 只列出第一个
		chainName := args[0]
		fmt.Println("\033[31;1m", chainName, "in", ruleChain[chainName], "with", f, "\033[0m")
		// 遍历每一个输出当前规则
		for _, j := range ruleChain[chainName] {
			goExecCmd(j, f, chainName)
		}
	}
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset all rules by chain",
}

var resetRuleCmd = &cobra.Command{
	Use:   "rule",
	Short: "reset all rules by chain",
	Run: func(cmd *cobra.Command, args []string) {
		doCmd(args, "-F")
	},
}

var resetCountCmd = &cobra.Command{
	Use:   "count",
	Short: "reset all package counts",
	Run: func(cmd *cobra.Command, args []string) {
		doCmd(args, "-Z")
	},
}

func goExecCmd(tableName string, flag string, ruleName string) {
	c := exec.Command("iptables", "-t", tableName, flag, ruleName)
	output, _ := c.CombinedOutput()
	if string(output)!=""{
		fmt.Println(string(output))
	}
}
