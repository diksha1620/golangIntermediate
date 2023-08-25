package main

import "fmt"

type employee interface {
	PrintName(name string)
	PrintSalary(sal int, tax int)
}

type Emp int

func (e Emp) PrintName(name string) {
	fmt.Println("id:", e)
	fmt.Println("name:", name)
}

func (e Emp) PrintSalary(sal int, tax int) int {
	salary := (sal * tax) / 100
	return salary - sal
}
func main() {
	fmt.Println("Hello Employee")
	e1 := Emp(1)
	e1.PrintName("diksha")
	fmt.Println("salary: ", e1.PrintSalary(12000, 300))
}
