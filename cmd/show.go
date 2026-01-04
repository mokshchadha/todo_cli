/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show all the existing todos",
	Long: `All the todos that are stored in side the csv file will appear to you in the format 
	ID, Description, status, createdAt`,
	Run: func(cmd *cobra.Command, args []string) {
		// here is the final logic that i want to execute
		// in this case read the data from the csv file
		file, err := os.Open("db.csv")
		if err != nil {
			log.Fatal("File db.csv not found")
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal("Could not read all from db.csv")
		}
		for _, record := range records {
			fmt.Println(record)
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
