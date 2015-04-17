package domains

import(

)

type CategoryRepository interface {
	FindById(id int) *Category
	Store(cate *Category)
}

type Category struct {
	Id int
	BookIds []int
}
