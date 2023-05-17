package data

type EmployeeService struct{}

// use pointer receiver is because we
// will adding a new employee
// which will be modifying the data
// inside employee service

type Employee struct {
	ID, Name, JobTitle string
}

func (es *EmployeeService) Add(e Employee) string {
	return ""
}
