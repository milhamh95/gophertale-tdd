package data_test

import (
	"gophertale-tdd/tdd/data"
	"testing"
)

// test is the first consumer of the api

func TestEmployeeService(t *testing.T) {
	t.Run("add new employee", func(t *testing.T) {
		e := data.Employee{
			ID:       "A-1",
			Name:     "Nikita",
			JobTitle: "The boss of everything",
		}

		es := data.EmployeeService{}
		id := es.Add(e)

		if id != e.ID {
			t.Fatalf("got (%v), want (%v)", id, e.ID)
		}
	})
}
