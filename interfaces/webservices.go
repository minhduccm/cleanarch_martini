package interfaces

import (
	"github.com/minhduccm/cleanarch_martini/domains"
	"encoding/json"
	"net/http"
	"github.com/go-martini/martini"
	"strconv"
)

type BookInteractor interface {
	GetBooksByCate(cateId int) []*domains.Book
	AddBook(bookId int, name string, price float64, cateId int)
}

type WebserviceHandler struct {
	BookInteractor BookInteractor
}

func (webserviceHandler *WebserviceHandler) ShowBooksByCate(params martini.Params, res http.ResponseWriter, req *http.Request) {
	//cateId := int(params["cateId"])
	cateId, err := strconv.Atoi(params["cateId"])
	var books []*domains.Book
	books = webserviceHandler.BookInteractor.GetBooksByCate(cateId)
	jsonData, err := json.Marshal(books)
	if err != nil {
	    http.Error(res, err.Error(), http.StatusInternalServerError)
	    return
	}

	//return jsonData
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)	
}