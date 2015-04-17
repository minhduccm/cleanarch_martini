package interfaces

import (
	"github.com/minhduccm/cleanarch_martini/domains"
)

type BookMapper struct {
	BookId int `bson:"book_id"` 
	Name string `bson:"book_name"`
	Price float64 `bson:"book_price"`
	Available bool `bson:"book_available"`
}

type CategoryMapper struct {
	CateId int `bson:"cate_id"` 
	BookIds []int `bson:"cate_bookids"`
}

type DbHandler interface {
	QueryBooksByIds(bookIds []int) []BookMapper
	QueryCateById(cateId int) CategoryMapper
	InsertBook(book BookMapper)
	InsertCate(cate CategoryMapper)
}

type DbBookRepo struct {
	DbHandler DbHandler
}


type DbCateRepo struct {
	DbHandler DbHandler
}


/* 
type DbRepo struct {
	dbHandler DbHandler
}

type DbBookRepo DbRepo
type DbCateRepo DbRepo
	
*/

func NewDbBookRepo(dbHandler DbHandler) *DbBookRepo {
	return &DbBookRepo{ DbHandler: dbHandler } 
}

func NewDbCateRepo(dbHandler DbHandler) *DbCateRepo {
	return &DbCateRepo{ DbHandler: dbHandler }
}

func (dbBookRepo *DbBookRepo) FindByIds(bookIds []int) []*domains.Book {
	rawBooks := dbBookRepo.DbHandler.QueryBooksByIds(bookIds)
	books := make([]*domains.Book, len(rawBooks))
	for i,book := range rawBooks {
		//books[i] = book.(*domains.Book) // wrong here as expect
		// convert BookModel => Book entity
		books[i] = &domains.Book{
			Id: book.BookId,
			Name: book.Name,
			Price: book.Price,
			Available: book.Available,
		}
	}
	return books
}

func (dbBookRepo *DbBookRepo) Store(book *domains.Book) {
	bookMapper := BookMapper{
		BookId: book.Id,
		Name: book.Name,
		Price: book.Price,
		Available: true,
	}
	dbBookRepo.DbHandler.InsertBook(bookMapper)
}

func (dbCateRepo *DbCateRepo) FindById(cateId int) *domains.Category {
	rawCate := dbCateRepo.DbHandler.QueryCateById(cateId)
	//return rawCate.(*domains.Category) // wrong here as expect
	return &domains.Category{
		Id: rawCate.CateId,
		BookIds: rawCate.BookIds,
	}
}

func (dbCateRepo *DbCateRepo) Store(cate *domains.Category) {
	cateMapper := CategoryMapper{
		CateId: cate.Id,
		BookIds: cate.BookIds,
	}
	dbCateRepo.DbHandler.InsertCate(cateMapper)
}
 
