package cli

import (
	"fmt"
	"log"

	"github.com/mohrezfadaei/passgen-cli/internal/generator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passgen-cli",
	Short: "Password and Key Generator CLI",
	Long:  "A command-line application to generate passwords and keys",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := generator.GeneratePassword()
		if err != nil {
			log.Fatalf("Error generating password: %v", err)
		}
		fmt.Println("Generated Password: ", password)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {}
