package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (R *Rectangle) Area() float64 {
	return R.width * R.height
}
func (R *Rectangle) Perimeter() float64 {
	return 2 * (R.width + R.height)
}

type Circle struct {
	r float64
}

func (C *Circle) Area() float64 {
	return math.Pi * C.r * C.r
}
func (C *Circle) Perimeter() float64 {
	return 2 * math.Pi * C.r
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeID int
	Person     Person
}

func (E *Employee) PrintInfo() {
	fmt.Printf("EmployeeID: %d, Name: %s, Age: %d\n", E.EmployeeID, E.Person.Name, E.Person.Age)
}

func main() {
	//1.周长和面积方法
	newCircle := &Circle{
		r: 2,
	}
	fmt.Println(newCircle.Area(), newCircle.Perimeter())
	newRectangle := &Rectangle{
		width:  2,
		height: 3,
	}
	fmt.Println(newRectangle.Area(), newRectangle.Perimeter())

	//2.组合的方式创建一个 Person 结构体
	em := &Employee{
		EmployeeID: 1,
		Person: Person{
			Name: "test",
			Age:  30,
		},
	}
	em.PrintInfo()

}
