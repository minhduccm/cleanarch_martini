package main

import(
	"fmt"
	"github.com/go-martini/martini"
)

func main() {
	fmt.Println("hello go on fucking ubuntu!")
	m := martini.Classic()

	m.Get("/", func() string {
		return "hello martini!"
	})

	m.Run()
}