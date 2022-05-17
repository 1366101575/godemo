package emp

type Manager struct {
}

func (m *Manager) GetTypeCode() string {
	return "manager"
}

func (m *Manager) PayAmount(employee *Employee) int {
	return employee.monthSalary() + employee.bonus()
}
