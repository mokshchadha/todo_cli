/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"todo-cli/utils"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new todo entry",
	Long:  `Your Todo entry should comprise of description and expected deadline`,
	Run: func(cmd *cobra.Command, args []string) {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the description of your todo")
		scanner.Scan()
		description := scanner.Text()
		fmt.Println("Created a todo for :", description)
		// update the csv
		createdAt := time.Now().UnixNano()
		existingLines, err := utils.CountLinesInCsv()
		if err != nil {
			log.Fatal(err)
		}
		id := strconv.FormatInt(int64(existingLines)+1, 10)
		record := []string{id, description, "incomplete", strconv.FormatInt(createdAt, 10)}

		file, err := os.OpenFile("db.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal("creation of file failed")
		}
		defer file.Close()
		w := csv.NewWriter(file)

		w.Write(record)
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
