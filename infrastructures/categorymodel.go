package infrastructures

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/minhduccm/cleanarch_martini/interfaces"
)

// type CategoryModel struct {
// 	CateId int `bson:"cate_id"` 
// 	BookIds []int `bson:"cate_bookids"`
// }

func (mongoDbHandler *MongoDbHandler) QueryCateById(cateId int) interfaces.CategoryMapper {
	sessionCopy := mongoDbHandler.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(Database).C(CategoriesCollection) 
	var cate interfaces.CategoryMapper
	collection.Find(bson.M{"cate_id": cateId}).One(&cate)
	return cate
}

func (mongoDbHandler *MongoDbHandler) InsertCate(cate interfaces.CategoryMapper) {
	// cateModel := CategoryModel{
	// 	CateId: cate.Id,
	// 	BookIds: cate.BookIds,
	// }
	sessionCopy := mongoDbHandler.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(Database).C(CategoriesCollection)
	err := collection.Insert(cate)
	if err != nil {
		fmt.Println("Insert cate error: ", err)
	} else {
		fmt.Println("Insert cate successfully")
	}
}