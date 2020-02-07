package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := &Saiyan{"Goku", 9000}
	goku.Super()
	fmt.Println(goku.Power)
}

// Super The type "*Saiyan" is the _receiver_ of the Super() method.
func (s *Saiyan) Super() {
	s = &Saiyan{"Gohan", 1000}
}

// NewSaiyan returns a Saiyan pointer. It's like a constructor, but a factory
// instead.
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}
