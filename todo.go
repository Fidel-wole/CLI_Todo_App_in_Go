package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{Title: title, Completed: false, CreatedAt: time.Now(), CompletedAt: nil}

	*todos = append(*todos, todo)

}

func (Todo *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*Todo) {
		err := errors.New("invalid index: %d")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) Complete(title string) {
	for i, todo := range *todos {
		if todo.Title == title {
			(*todos)[i].Completed = true
			(*todos)[i].CompletedAt = &time.Time{}
		}
	}
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if !isCompleted{
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
    t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) List() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for i, todo := range *todos {
		completed := "❌"
		if todo.Completed {
			completed = "✅"
		}
		completedAt := ""
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}
		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
