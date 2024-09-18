package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (td *Todo) String() string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(strconv.Itoa(td.Id))
	sb.WriteString(") ")
	sb.WriteString(td.Title)
	sb.WriteString(": ")
	if td.Completed {
		sb.WriteString("✅")
	} else {
		sb.WriteString("⬜")
	}

	return sb.String()
}

func Load(filename string) (Todos, error) {
	var todos Todos

	jsonFile, err := os.Open(filename)
	if err != nil {
		return todos, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return todos, err
	}

	if len(byteValue) == 0 {
		return todos, nil
	}

	err = json.Unmarshal(byteValue, &todos)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (tds Todos) Save(filename string) error {
	jsonData, err := json.MarshalIndent(tds, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (tds Todos) String() string {
	var sb strings.Builder

	for _, t := range tds.Todos {
		sb.WriteString(t.String())
		sb.WriteString("\n")
	}

	return sb.String()
}

func (tds Todos) NextId() int {
	highestId := -1
	for _, t := range tds.Todos {
		if t.Id > highestId {
			highestId = t.Id
		}
	}
	return highestId + 1
}

func (tds Todos) Add(todo Todo) Todos {
	tds.Todos = append(tds.Todos, todo)
	return tds
}

func (tds Todos) Complete(id int) (Todos, error) {
	for i := range tds.Todos {
		if tds.Todos[i].Id == id {
			tds.Todos[i].Completed = true
			return tds, nil
		}
	}

	return tds, fmt.Errorf("todos with id %d could not be found", id)
}

func (tds Todos) Delete(id int) (Todos, error) {
	for i := range tds.Todos {
		if tds.Todos[i].Id == id {
			tds.Todos = append(tds.Todos[:i], tds.Todos[i+1:]...)
			return tds, nil
		}
	}

	return tds, fmt.Errorf("todos with id %d could not be found", id)
}
