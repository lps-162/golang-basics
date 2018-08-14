package main

import (
	"fmt"
)

// Namer interface
type Namer interface {
	Name() string
}

// Greet method of Namer interface
func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}

// User structure
type User struct {
	FirstName, LastName string
}

// Name method, implementing Name method
func (u *User) Name() string {
	return fmt.Sprintf("User %s %s", u.FirstName, u.LastName)
}

// Customer type and its implementation methods
type Customer struct {
	ID       int
	FullName string
}

// Name ...implementing Namer() interface
func (c *Customer) Name() string {
	return fmt.Sprintf("Customer %s", c.FullName)
}

func main() {
	u := &User{"Shivan", "LP"}
	fmt.Println(Greet(u))

	c := &Customer{42, "Servarayan Ishan"}
	fmt.Println(Greet(c))
}
