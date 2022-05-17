package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := make(map[string]int, 8)
	m["北京"] = 99
	m["深圳"]++
	fmt.Printf("%#v\n", m)

	_, ok := m["北京"]
	if ok {
		delete(m, "北京")
	}
	fmt.Printf("%#v\n", m)

	fmt.Println(http.StatusOK)

}
