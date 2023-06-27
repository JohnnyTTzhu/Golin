package cmd

import (
	"github.com/spf13/cobra"
	"golin/domain"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "基于DNS碰撞子域名",
	Run:   domain.ParseFlags,
}

func init() {
	rootCmd.AddCommand(domainCmd)
	domainCmd.Flags().StringP("url", "u", "", "此参数是指定扫描域名")
	domainCmd.Flags().StringP("file", "f", "", "此参数是指定字典文件")
	domainCmd.Flags().IntP("chan", "c", 30, "并发数量")
}
