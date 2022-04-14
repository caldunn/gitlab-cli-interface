package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	config2 "gitlab-cli-interface/config"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "NOT IMPLEMENTED",
	Long:  `NOT IMPLEMENTED`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config2.GetConfig()

		remarsh, _ := yaml.Marshal(config)
		fmt.Println(string(remarsh))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
