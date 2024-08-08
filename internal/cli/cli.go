package cli

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/mohrezfadaei/passgen-cli/internal/generator"
	"github.com/spf13/cobra"
)

var length int
var excludeUppercase bool
var excludeLowercase bool
var excludeDigits bool
var excludeSymbols bool
var includeChars string
var copyToClipboard bool

var rootCmd = &cobra.Command{
	Use:   "passgen-cli",
	Short: "Password and Key Generator CLI",
	Long:  "A command-line application to generate passwords and keys",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := generator.GeneratePassword(length, excludeUppercase, excludeLowercase, excludeDigits, excludeSymbols, includeChars)
		if err != nil {
			log.Fatalf("Error generating password: %v", err)
		}
		fmt.Println(password)
		if copyToClipboard {
			err := clipboard.WriteAll(password)
			if err != nil {
				log.Fatalf("Error copying to clipboard: %v", err)
			}
			fmt.Println("Password copied to clipboard")
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().IntVarP(&length, "length", "l", 16, "Length of the password")
	rootCmd.Flags().BoolVar(&excludeUppercase, "no-uppercase", false, "Exclude uppercase letters")
	rootCmd.Flags().BoolVar(&excludeLowercase, "no-lowercase", false, "Exclude lowercase letters")
	rootCmd.Flags().BoolVar(&excludeDigits, "no-digits", false, "Exclude digits")
	rootCmd.Flags().BoolVar(&excludeSymbols, "no-symbols", false, "Exclude symbols")
	rootCmd.Flags().StringVar(&includeChars, "include-chars", "", "Include specific characters (incl., $@#$%^&*()_+{}|:\"<>?,./~`)")
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "c", false, "Copy password to clipboard")
}
