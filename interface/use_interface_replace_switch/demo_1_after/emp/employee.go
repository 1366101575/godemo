package emp

// Employee 员工 基础月薪
type Employee struct {
	Name         string
	EmployeeType EmployeeType
}

func (e *Employee) monthSalary() int {
	return 3000
}

func (e *Employee) allowance() int {
	return 1000
}

func (e *Employee) bonus() int {
	return 2000
}
