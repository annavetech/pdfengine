// Copyright 2026 Anna Veretennykova
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/annavetech/pdfengine/internal/engine"
)

var rootCmd = &cobra.Command{
	Use:   "pdfengine",
	Short: "Convert documents to PDF",
	Long:  "PDF Engine: convert documents to PDF from the command line.",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("PDF Engine v%s\n", engine.EngineVersion)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(pdfCmd)
}
