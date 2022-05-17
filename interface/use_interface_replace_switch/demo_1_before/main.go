package main

import (
	"fmt"
	"godemo/interface/use_interface_replace_switch/demo_1_before/emp"
)

func main() {
	e := emp.Employee{Name: "张三", TypeCode: "engineer"}
	amount, err := e.PayAmount()
	if err != nil {
		fmt.Println("err: " + err.Error())
		return
	}
	fmt.Printf("name:%s type:%s amout:%d\n", e.Name, e.TypeCode, amount)

	e = emp.Employee{Name: "李四", TypeCode: "salesman"}
	amount, err = e.PayAmount()
	if err != nil {
		fmt.Println("err: " + err.Error())
		return
	}
	fmt.Printf("name:%s type:%s amout:%d\n", e.Name, e.TypeCode, amount)

	e = emp.Employee{Name: "王五", TypeCode: "manager"}
	amount, err = e.PayAmount()
	if err != nil {
		fmt.Println("err: " + err.Error())
		return
	}
	fmt.Printf("name:%s type:%s amout:%d\n", e.Name, e.TypeCode, amount)
}
