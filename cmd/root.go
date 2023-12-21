/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/carantes/wc-go/wc"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wc-go",
	Short: "This is my own version of the wc command",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		lines, _ := cmd.Flags().GetBool("lines")
		words, _ := cmd.Flags().GetBool("words")
		characters, _ := cmd.Flags().GetBool("characters")

		var fileContent []byte

		// Read from file
		if len(args) > 0 {
			var err error
			fileContent, err = os.ReadFile(args[0])
			if err != nil {
				fmt.Printf("Error reading file: %v", err)
				os.Exit(1)
			}
		} else {
			// Read from stdin
			var err error
			fileContent, err = io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("No filename or stdin provided")
				os.Exit(1)
			}
		}

		var result int
		if lines {
			result = wc.CountLines(fileContent)
		} else if words {
			result = wc.CountWords(fileContent)
		} else if characters {
			result = wc.CountCharacters(fileContent)
		} else {
			result = wc.CountBytes(fileContent)
		}

		if len(args) == 0 {
			fmt.Printf("%d\n", result)
		} else {
			fmt.Printf("%d %s\n", result, args[0])
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("bytes", "c", false, "Number of bytes")
	rootCmd.PersistentFlags().BoolP("lines", "l", false, "Number of lines")
	rootCmd.PersistentFlags().BoolP("words", "w", false, "Number of words")
	rootCmd.PersistentFlags().BoolP("characters", "m", false, "Number of characters")
}
