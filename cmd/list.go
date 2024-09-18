package cmd

import (
	"fmt"
	"github.com/Kuzat/go-cli/todo"
	"github.com/spf13/cobra"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the todo list",
	Run: func(cmd *cobra.Command, args []string) {
		filename := "todos.json"
		todos, err := todo.Load(filename)
		if err != nil {
			log.Fatalf("Error loading file %s: %v", filename, err)
		}

		fmt.Println(todos.String())
	},
}
