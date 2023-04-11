package main

import (
	"errors"
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
