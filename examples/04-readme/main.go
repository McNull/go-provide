package main

import (
	"fmt"

	"github.com/mcnull/go-provide"
)

////////////////////////
// Setup DI
////////////////////////

func init() {
	provide.Set[MyRepo](NewMyRepo)
	provide.Set[MyService](NewMyService)
}

func main() {
	myService, err := provide.Get[MyService]()

	if err != nil {
		panic(err)
	}

	myService.Print() // Output: Hello, World!
}

////////////////////////
// Repository
////////////////////////

type MyRepo interface {
	Get() string
}

type myRepo struct{}

func (r *myRepo) Get() string {
	return "Hello, World!"
}

func NewMyRepo( /* add dependencies here*/ ) *myRepo {
	return &myRepo{}
}

////////////////////////
// Service
////////////////////////

type MyService interface {
	Print()
}

type myService struct {
	Repo MyRepo
}

func (s *myService) Print() {
	fmt.Println(s.Repo.Get())
}

func NewMyService(myRepo MyRepo /* add dependencies here */) *myService {
	return &myService{
		Repo: myRepo,
	}
}
