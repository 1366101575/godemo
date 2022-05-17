package emp

type Engineer struct {
}

func (e *Engineer) GetTypeCode() string {
	return "engineer"
}

func (e *Engineer) PayAmount(employee *Employee) int {
	return employee.monthSalary()
}
