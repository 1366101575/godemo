package emp

import "errors"

// Employee 员工 基础月薪
type Employee struct {
	Name     string
	TypeCode string
}

func (e *Employee) PayAmount() (int, error) {
	switch e.TypeCode {
	case "engineer":
		return e.monthSalary(), nil
	case "salesman":
		// 月薪 + 补贴
		return e.monthSalary() + e.allowance(), nil
	case "manager":
		// 月薪 + 奖金
		return e.monthSalary() + e.bonus(), nil
	default:
		return 0, errors.New("type not support")
	}
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
