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

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		data: make(map[string]Employee),
	}
}

// use *string in case something goes wrong
// we want to make it very clear to whoever
// call add, can't use this value.
// if we use string, the caller can use the value.
// alternative, caller need to check id is empty or not
func (es *EmployeeService) Add(e Employee) (*string, error) {
	es.data[e.ID] = e
	return &e.ID, nil
}
