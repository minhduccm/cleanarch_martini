package usecases

import (
	"github.com/minhduccm/cleanarch_martini/domains"
)

type BookInteractor struct {
	BookRepository domains.BookRepository
	CateRepository domains.CategoryRepository
}

func (bookInteractor *BookInteractor) GetBooksByCate(cateId int) []*domains.Book {
	var books []*domains.Book
	cate := bookInteractor.CateRepository.FindById(cateId)
	if cate != nil {
		bookIds := cate.BookIds
		books = bookInteractor.BookRepository.FindByIds(bookIds)
	}
	return books
}

func (bookInteractor *BookInteractor) AddBook(bookId int, name string, price float64, cateId int) {
	book := &domains.Book{
		Id: bookId,
		Name: name,
		Price: price,		
	}
	// store new book to Books collection
	bookInteractor.BookRepository.Store(book)
	// push bookid to cate that it belongs to
	cate := bookInteractor.CateRepository.FindById(cateId)
	if cate != nil {
		cate.BookIds = append(cate.BookIds, book.Id)
		bookInteractor.CateRepository.Store(cate)	
	} else {
		newCate := &domains.Category{
			Id: cateId,
			BookIds: []int{bookId},
		}
		bookInteractor.CateRepository.Store(newCate)
	}
	
}