package main

import(
	//"fmt"
	"github.com/go-martini/martini"
	"github.com/minhduccm/cleanarch_martini/interfaces"
	"github.com/minhduccm/cleanarch_martini/infrastructures"
	"github.com/minhduccm/cleanarch_martini/usecases"	
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {		
		return "Hello martini!"
	})

	mongoHandler := infrastructures.NewMongoDbHandler()
	dbBookRepo := interfaces.NewDbBookRepo(mongoHandler)
	dbCateRepo := interfaces.NewDbCateRepo(mongoHandler)
	bookInteractor := &usecases.BookInteractor{
		BookRepository: dbBookRepo,
		CateRepository: dbCateRepo,
	}

	webserviceHandler := &interfaces.WebserviceHandler{
		BookInteractor: bookInteractor,
	}

	// m.Get("/books/:cateId", func() {
	// 	return webserviceHandler.ShowBooksByCate
	// })

	m.Get("/books/cate/:cateId", webserviceHandler.ShowBooksByCate)

	m.Run()
}