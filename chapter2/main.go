package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	fmt.Println(goku.Power)
}
