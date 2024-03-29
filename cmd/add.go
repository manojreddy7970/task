package cmd

import ("fmt"
	"github.com/Manoj-Reddy-Vadde/task/db"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		// for i,arg := range args {
		// 	fmt.Println(i, ":", arg)
		// }
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init()  {
	RootCmd.AddCommand(addCmd)	
}