package main

import (
	"fmt"
	"godemo/interface/use_interface_replace_switch/demo_1_after/emp"
)

// 员工与薪资

func main() {
	e := emp.Employee{
		Name:         "张三",
		EmployeeType: &emp.Engineer{},
	}
	fmt.Printf("name:%s type:%s amount:%d\n", e.Name, e.EmployeeType.GetTypeCode(), e.EmployeeType.PayAmount(&e))

	e = emp.Employee{
		Name:         "李四",
		EmployeeType: &emp.Salesman{},
	}
	fmt.Printf("name:%s type:%s amount:%d\n", e.Name, e.EmployeeType.GetTypeCode(), e.EmployeeType.PayAmount(&e))

	e = emp.Employee{
		Name:         "王五",
		EmployeeType: &emp.Manager{},
	}
	fmt.Printf("name:%s type:%s amount:%d\n", e.Name, e.EmployeeType.GetTypeCode(), e.EmployeeType.PayAmount(&e))
}
