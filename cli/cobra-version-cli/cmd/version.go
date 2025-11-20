package cmd

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	GitCommit  = "unknown"
	GitVersion = "unknown"
	BuildDate  = "unknown"
	RepoUrl    = "unknown"
)

var name = `
                   _                   ___
 _  _____ _______ (_)__  ___  ________/ (_)
| |/ / -_) __(_-</ / _ \/ _ \/___/ __/ / /
|___/\__/_/ /___/_/\___/_//_/    \__/_/_/
`

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show version",
	Example: "cobra-version-cli version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s%s\n\n", name, RepoUrl)

		type KV struct {
			K string
			V string
		}
		kvs := []KV{
			{K: "GitVersion", V: GitVersion},
			{K: "GitCommit", V: GitCommit},
			{K: "BuildDate", V: BuildDate},
			{K: "GoVersion", V: runtime.Version()},
			{K: "Compiler", V: runtime.Compiler},
			{K: "Platform", V: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)},
		}

		tw := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
		defer tw.Flush()
		for _, kv := range kvs {
			fmt.Fprintf(tw, "%s:\t%s\n", kv.K, kv.V)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
