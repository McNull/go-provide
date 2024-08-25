package main

import (
	"fmt"

	"github.com/mcnull/go-provide"
)

type MyService struct{}

func (s *MyService) DoSomething() {
	fmt.Println("Doing something")
}

type MyServiceInterface interface {
	DoSomething()
}

func main() {

	// create an instance of MyService and provide it by using an interface
	s := new(MyService)
	provide.Set[MyServiceInterface](s)

	// ... do something else
	someTimeLater()
}

func someTimeLater() {
	// get the instance of MyService by using the interface
	s, err := provide.Get[MyServiceInterface]()

	if err != nil {
		fmt.Println(err)
		return
	}

	// call the method on the instance
	s.DoSomething()
}
