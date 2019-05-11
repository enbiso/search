package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	searchCmd := searchCmd()
	searchCmd.AddCommand(versionCmd())

	if err := searchCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func searchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "search",
		Run: func(cmd *cobra.Command, args []string) {
			searchInit(
				cmd.Flags().Lookup("data").Value.String(),
				cmd.Flags().Lookup("addr").Value.String(),
			)
		},
	}
	cmd.Flags().StringP("data", "d", "./data", "Data directory (./data)")
	cmd.Flags().StringP("addr", "a", ":8080", "Listen address (:8080)")
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("data", cmd.Flags().Lookup("data"))
	return cmd
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of enbiso search",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Enbiso Search v0.1.0")
		},
	}
}
