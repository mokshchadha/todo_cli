/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo-cli/utils"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a todo item as done",
	Long:  `Mark a todo item as done by providing its ID`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		fmt.Println("Marking todo as done:", id)
		utils.MarkTodoAsCompleted(id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.Flags().StringP("id", "i", "", "Todo ID to mark as done (required)")
	doneCmd.MarkFlagRequired("id")
}
