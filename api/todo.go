package main

import "time"

// TodoAPI type
type TodoAPI struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// TodosAPI - array of Todo
type TodosAPI []TodoAPI
