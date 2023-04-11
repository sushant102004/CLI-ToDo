package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type ToDo struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// This will contain a list of ToDo items.
type ToDoList []ToDo

// Reciever function to create new ToDo and append to ToDoList
func (t *ToDoList) addTask(task string) {
	newTask := ToDo{
		Title:       task,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, newTask)

}

func (t *ToDoList) completedTask(taskIndex int) error {
	list := *t

	if taskIndex <= 0 || taskIndex >= len(list) {
		return errors.New("invalid task index")
	}

	list[taskIndex-1].CompletedAt = time.Now()
	list[taskIndex].IsCompleted = true

	return nil
}

func (t *ToDoList) deleteTask(index int) error {
	list := *t

	if index <= 0 || index >= len(list) {
		return errors.New("invalid index")
	}

	*t = append(list[:index-1], list[index+1:]...)

	return nil
}

func (t *ToDoList) loadFile(fName string) error {
	file, err := os.ReadFile(fName)

	if err != nil {
		// ErrNotExist will check if file does not exist.
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		panic("file empty!!")
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil
}
