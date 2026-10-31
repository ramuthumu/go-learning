package main

import (
	"errors"
	"fmt"
)

// Task represents a single to-do item (Lesson 8: Structs)
type Task struct {
	ID    int
	Title string
	Done  bool
}

// TaskList manages a collection of tasks
type TaskList struct {
	Tasks  []Task      // Lesson 5: Slices
	ByID   map[int]int // Lesson 6: Maps (Task ID -> Slice index)
	NextID int
}

// NewTaskList initializes a fresh TaskList (Lesson 7: Functions)
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  []Task{},
		ByID:   make(map[int]int),
		NextID: 1,
	}
}

// Add creates a new task (Pointer receiver so it modifies the list)
func (tl *TaskList) Add(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("task title cannot be empty")
	}

	task := Task{
		ID:    tl.NextID,
		Title: title,
		Done:  false,
	}

	tl.ByID[task.ID] = len(tl.Tasks)
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	return task, nil
}

// Complete marks a task as done by ID
func (tl *TaskList) Complete(id int) error {
	// Comma-OK idiom to check if ID exists in map
	idx, ok := tl.ByID[id]
	if !ok {
		return fmt.Errorf("task #%d not found", id)
	}
	tl.Tasks[idx].Done = true
	return nil
}

// PrintAll displays all tasks using a for-range loop
func (tl TaskList) PrintAll() {
	fmt.Println("\n=== My Go Learning Tasks ===")
	for _, t := range tl.Tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%s #%d: %s\n", status, t.ID, t.Title)
	}
}

func main() {
	list := NewTaskList()

	// Add tasks
	_, _ = list.Add("Learn Go variables and types")
	_, _ = list.Add("Master slices, maps, and loops")
	_, _ = list.Add("Understand structs and pointers")

	// Try adding an invalid empty task to test error handling
	if _, err := list.Add(""); err != nil {
		fmt.Println("Expected validation error:", err)
	}

	// Mark tasks 1 and 2 complete
	_ = list.Complete(1)
	_ = list.Complete(2)

	// Print summary
	list.PrintAll()
}
