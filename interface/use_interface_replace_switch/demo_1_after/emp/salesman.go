package emp

type Salesman struct {
}

func (s *Salesman) GetTypeCode() string {
	return "salesman"
}

func (s *Salesman) PayAmount(employee *Employee) int {
	return employee.monthSalary() + employee.allowance()
}
