package cmd

import (
	"fmt"
	"github.com/Kuzat/todo-cli/todo"
	"github.com/spf13/cobra"
	"log"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task to the todo list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := "todos.json"
		title := args[0]
		todos, err := todo.Load(filename)
		if err != nil {
			log.Fatalf("Error loading file %s: %v", filename, err)
		}

		todo := todo.Todo{
			Id:        todos.NextId(),
			Title:     title,
			Completed: false,
		}

		todos = todos.Add(todo)

		fmt.Printf("Task added: %s\n", todo.String())

		err = todos.Save(filename)
		if err != nil {
			log.Fatalf("Error saving todo file, %v", err)
		}
	},
}
