package cmd

import (
	"github.com/spf13/cobra"
	"gitlab-cli-interface/http"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Output your issues from gitlab specified in YAML",
	Long:  `Output your issues from gitlab specified in YAML`,
	Run: func(cmd *cobra.Command, args []string) {
		http.AllIssues()
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
