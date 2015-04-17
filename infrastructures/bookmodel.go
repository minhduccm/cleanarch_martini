package infrastructures

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/minhduccm/cleanarch_martini/interfaces"
)

// type BookModel struct {
// 	BookId int `bson:"book_id"` 
// 	Name string `bson:"book_name"`
// 	Price float64 `bson:"book_price"`
// 	Available bool `bson:"book_available"`
// }

func (mongoDbHandler *MongoDbHandler) QueryBooksByIds(bookIds []int) []interfaces.BookMapper {
	sessionCopy := mongoDbHandler.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(Database).C(BooksCollection) 
	var books []interfaces.BookMapper
	collection.Find(bson.M{"book_id": bson.M{"$in": bookIds}}).All(&books)
	return books
}


func (mongoDbHandler *MongoDbHandler) InsertBook(book interfaces.BookMapper) {
	// bookModel := BookModel{
	// 	BookId: book.Id,
	// 	Name: book.Name,
	// 	Price: book.Price,
	// 	Available: true,
	// }
	sessionCopy := mongoDbHandler.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(Database).C(BooksCollection)
	err := collection.Insert(book)
	if err != nil {
		fmt.Println("Insert book error: ", err)
	} else {
		fmt.Println("Insert book successfully")
	}
}

