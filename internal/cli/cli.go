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
var justInclude string
var copyToClipboard bool

var rootCmd = &cobra.Command{
	Use:   "passgen-cli",
	Short: "Password and Key Generator CLI",
	Long:  "A command-line application to generate passwords and keys",
	Run: func(cmd *cobra.Command, args []string) {
		if includeChars != "" && justInclude != "" {
			log.Fatalf("Flags --include-chars and --just-include cannot be used together")
		}
		if (excludeUppercase || excludeLowercase || excludeDigits || excludeSymbols) && justInclude != "" {
			log.Fatalf("Flags --exclude-uppercase, --exclude-lowercase, --exclude-digits, and --exclude-symbols cannot be used together with --just-include")
		}

		password, err := generator.GeneratePassword(length, excludeUppercase, excludeLowercase, excludeDigits, excludeSymbols, includeChars, justInclude)
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
	rootCmd.Flags().StringVar(&justInclude, "just-include", "", "Just include specific characters (excl., @#$%^&)")
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "c", false, "Copy password to clipboard")
}
