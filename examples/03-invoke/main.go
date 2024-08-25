package main

import (
	"fmt"

	"github.com/mcnull/go-provide"
)

type serviceData string

type myService struct {
	data serviceData
}

func (s *myService) Print() {
	fmt.Println(s.data)
}

type MyService interface {
	Print()
}

func NewMyService(data serviceData) *myService {
	return &myService{
		data: data,
	}
}

func init() {
	sd := serviceData("Hello, World!")
	provide.Set[serviceData](sd)
	provide.Set[MyService](NewMyService)
}

func invokeMe(s MyService) {
	s.Print()
}

func main() {
	provide.Invoke(invokeMe)
}
