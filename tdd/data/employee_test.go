package data_test

import (
	"errors"
	"gophertale-tdd/tdd/data"
	"testing"
)

// test is the first consumer of the api

func TestEmployeeService(t *testing.T) {
	// es := data.EmployeeService{}
	// we can put es in the beginning but the test
	// will not be isolated

	t.Run("add new employee", func(t *testing.T) {
		e := data.Employee{
			ID:       "A-1",
			Name:     "Nikita",
			JobTitle: "The boss of everything",
		}

		es := data.NewEmployeeService()
		id, err := es.Add(e)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if *id != e.ID {
			t.Fatalf("got (%v), want (%v)", *id, e.ID)
		}
	})

	t.Run("add empty name employee", func(t *testing.T) {
		e := data.Employee{
			ID:       "A-1",
			JobTitle: "The boss of everything",
		}

		es := data.NewEmployeeService()
		id, err := es.Add(e)
		if id != nil {
			t.Fatalf("unexpected non nil id: %v", *id)
		}

		expectedErr := errors.New("empty name")
		if err == nil || err.Error() != expectedErr.Error() {
			t.Fatalf("got (%v), want (%v)", err, expectedErr)
		}
	})
}
