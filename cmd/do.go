package cmd

import (
	"fmt"
	"strconv"
	
	"github.com/Manoj-Reddy-Vadde/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task has complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _,arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed o paese the arguement;")
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <=0 || id > len(tasks) {
				fmt.Println("Invalid  task number: ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)

			if err!= nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id ,err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)

			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
