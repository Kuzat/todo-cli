package cmd

import (
	"fmt"
	"github.com/Kuzat/go-cli/todo"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete todo task with the given id",
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

		todos, err = todos.Delete(id)
		if err != nil {
			log.Fatalf("Error deleting task: %v", err)
		}

		fmt.Printf("Deleted task (%d)\n", id)
		fmt.Println(todos.String())

		err = todos.Save(filename)
		if err != nil {
			log.Fatalf("Error saving todo file, %v", err)
		}
	},
}
