package data

import (
	"os"
	"testing"
)

var todos = New()

func TestMain(m *testing.M) {
	todos.Add(Item{Title: "buy something", Description: "NOW", Done: false})
	todos.Add(Item{Title: "walk the dog", Description: "grab the pet and walk it arount the yard", Done: false})
	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	if len(todos.Items) == 0 {
		t.Errorf("Add() failed. Expected: %d, Got: %d", len(todos.Items), 0)
	}
}

func TestGetAll(t *testing.T) {
	result := todos.GetAll()

	if len(result) == 0 {
		t.Errorf("GetAll() failed. Expected: %d, Got: %d", len(todos.Items), 0)
	}
}
