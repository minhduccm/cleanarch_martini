package domains

import(

)

type BookRepository interface {
	FindByIds(ids []int) []*Book
	Store(book *Book)
}

type Book struct {
	Id int
	Name string
	Price float64
	Available bool
}
