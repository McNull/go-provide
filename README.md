# go-provide

`go-provide` is a simple dependency injection module for Go. It allows you to define a set of providers that can be used to create instances of objects. 

## Installation

To install the module, run:

```sh
go get github.com/mcnull/go-provide
```

## Usage

```go
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
```

## Api documentation

### Set
Sets a provider for a type. 

The provider can be a (factory) function or an instance of the type. If the provider is a function, all arguments of the function are resolved and passed to the function. The return value of the function is used as the instance of the type.

```go
// Set a custom int type by value

type myInt int
provide.Set[myInt](myInt(42))

// Set a custom struct type by factory function

type myStruct struct {
  Number myInt
}

NewMyStruct := func(number myInt) *myStruct {
  return &myStruct{
    Number: number // injected
  }
}

provide.Set[*myStruct](NewMyStruct)
```

### Get
Resolves an instance of the provided type.

Returns an earlier registered instance of a type, with all dependencies resolved.

```go
// Get the custom struct type
instance, err := provide.Get[*myStruct]()

if err != nil {
  panic(err)
}

println(instance.Number) // Output: 42
```

### Invoke
Executes a function with all dependencies resolved.

```go
function myFunction(myStruct *myStruct) {
  println(myStruct.Number)
}

provide.Invoke(myFunction)
```

## License
MIT