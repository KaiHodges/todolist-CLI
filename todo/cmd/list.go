/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
 		rows := db.ListTodo()
		defer rows.Close()
		
		var ids []int 
		var tasks []string 
		var dates []string 
		var bools []int

		for rows.Next() {
			var id int 
			var task string 
			var date string 
			var bool int 
			if err := rows.Scan(&id, &task, &date, &bool); err != nil {
				log.Fatal(err)
			}
			ids = append(ids, id)
			tasks = append(tasks, task)
			dates = append(dates, date)
			bools = append(bools, bool)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		for i := range ids {
			fmt.Printf("%d, %s, %s, %t\n", i, tasks[i], dates[i], (bools[i] != 0))
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
