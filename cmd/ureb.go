/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// urebCmd represents the ureb command
var urebCmd = &cobra.Command{
	Use:   "ureb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		flagValue, _ := cmd.Flags().GetString("one")

		fmt.Printf("flagValue: %v\n", flagValue)

		// if flagValue != "" {
		// 	fmt.Println("ureb is awesome")
		// } else {
		// 	fmt.Println("ureb is awesome - one")
		// }
	},
}

func init() {
	rootCmd.AddCommand(urebCmd)

	urebCmd.PersistentFlags().String("one", "", "this prints one after normal thing")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// urebCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// urebCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
