package main

import (
	"fmt"

	"github.com/mcnull/go-provide"
	"github.com/mcnull/go-provide/examples/02-injection/pkg"
)

// setup the dependencies
func init() {
	// some default data to work with
	users := []pkg.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	provide.Set[[]pkg.User](users)
	provide.Set[pkg.UserRepository](pkg.NewUserRepository)
	provide.Set[pkg.UserService](pkg.NewUserService)
}

func main() {
	// get the instance of UserService by using the interface
	s, err := provide.Get[pkg.UserService]()

	if err != nil {
		fmt.Println(err)
		return
	}

	// call the method on the instance
	fmt.Println(s.PrintAll())
}
