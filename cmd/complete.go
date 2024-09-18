package cmd

import (
	"fmt"
	"github.com/Kuzat/todo-cli/todo"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Complete a task with the given id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Given id %s is not a valid int id. %v", args[0], err)
		}
		filename := "todos.json"
		todos, err := todo.Load(filename)
		if err != nil {
			log.Fatalf("Error loading file %s: %v", filename, err)
		}

		todos, err = todos.Complete(id)
		if err != nil {
			log.Fatalf("Error completing task: %v", err)
		}

		fmt.Printf("Completed task (%d)\n", id)
		fmt.Println(todos.String())

		err = todos.Save(filename)
		if err != nil {
			log.Fatalf("Error saving todo file, %v", err)
		}
	},
}
