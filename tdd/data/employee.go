package data

type EmployeeService struct {
	data map[string]Employee
}

// use pointer receiver is because we
// will adding a new employee
// which will be modifying the data
// inside employee service

type Employee struct {
	ID, Name, JobTitle string
}

func (es *EmployeeService) Add(e Employee) string {
	es.data[e.ID] = e
	return e.ID
}
