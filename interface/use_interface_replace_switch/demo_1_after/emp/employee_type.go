package emp

type EmployeeType interface {
	GetTypeCode() string
	PayAmount(employee *Employee) int
}
