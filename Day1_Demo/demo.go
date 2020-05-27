package main
import "fmt"
type calculateSalary interface {
	computeSalary() int
}
type pEmployee struct {
	 eid int
	 salary int
	 pf int
}
type cEmployee struct {
	eid int
	salary int
}
type fEmployee struct {
	eid int
	salary int
	hour int
}
func (p pEmployee) computeSalary() int{
	return p.salary+p.pf
}
func (c cEmployee) computeSalary() int{
	return c.salary
}
func (f fEmployee) computeSalary() int{
	return f.salary * f.hour
}
func totalExpense(a []calculateSalary) int {
	//fmt.Println("totak expense is",pSal+cSal)
	var sum int
	for _,value := range a{
		sum += value.computeSalary()
	}
	return sum
}
func main() {
	p := pEmployee{100,10000,2000}
	c := cEmployee{201,5000}
	f := fEmployee{301,3000,2}
	//perSalary := p.computeSalary()
	//contSalary := c.computeSalary()
	exp := []calculateSalary{p,c,f}
	fmt.Println("expense slice is",exp)
	expense := totalExpense(exp)
	fmt.Println("totak expense",expense);
	
}