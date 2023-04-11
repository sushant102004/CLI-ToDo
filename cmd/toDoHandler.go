package main

import "time"

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
